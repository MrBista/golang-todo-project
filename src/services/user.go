package services

import (
	"context"
	"database/sql"

	"github.com/MrBista/golang-todo-project/helper"
	"github.com/MrBista/golang-todo-project/src/dto/request"
	"github.com/MrBista/golang-todo-project/src/dto/response"
	"github.com/MrBista/golang-todo-project/src/model"
	"github.com/MrBista/golang-todo-project/src/repository"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	RegisterUser(ctx context.Context, req request.RegisterUserRequest) response.RegisterUserResponse
	LoginUser(ctx context.Context, req request.LoginUserReq) response.LoginUserRes
}

type UserServiceImpl struct {
	userRepository repository.UserRepositry
	DB             *sql.DB
}

func NewUserService(userRepo repository.UserRepositry, db *sql.DB) UserService {
	return &UserServiceImpl{
		userRepository: userRepo,
		DB:             db,
	}
}

func (s *UserServiceImpl) RegisterUser(ctx context.Context, reqBody request.RegisterUserRequest) response.RegisterUserResponse {
	trx, err := s.DB.Begin()

	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	defer func() {
		err := recover()
		if err != nil {
			logrus.Errorf("Error %v", err)
			errorRolback := trx.Rollback()
			if errorRolback != nil {
				panic(err)
			}
		} else {
			errorCommit := trx.Commit()
			if errorCommit != nil {
				panic(errorCommit)
			}
		}
	}()

	password, err := helper.HashPassword(reqBody.Password)

	if err != nil {
		// bisa pakai pesan custom
		panic(err)
	}

	user := model.User{
		Username: reqBody.Username,
		Email:    reqBody.Email,
		Password: string(password), // sementara plan text
		FullName: reqBody.FullName,
	}

	userSaved := s.userRepository.CreateUser(ctx, trx, user)

	userResponse := response.RegisterUserResponse{
		Id:       userSaved.Id,
		Email:    userSaved.Email,
		Username: userSaved.Username,
		FullName: userSaved.FullName,
	}

	return userResponse
}

func (s *UserServiceImpl) LoginUser(ctx context.Context, req request.LoginUserReq) response.LoginUserRes {

	// 1. cari dulu user nya
	// 2. kalau ga ada maka panic
	// 3. kala ada maka check dulu passwordnya
	// 4. kalau ga match maka panic
	// 5. genereate jwt dengan user id dan exp time dari viper

	trx, err := s.DB.Begin()

	if err != nil {
		helper.Logger().Error(err)
		panic(err)
	}

	findUserIdentifier, err := s.userRepository.FindByEmailOrUsername(ctx, trx, req.Identifier)

	if err != nil {
		helper.Logger().Error(err)
		panic(err)
	}

	if err := helper.ComparePassword(req.Password, findUserIdentifier.Password); err != nil {
		panic("username/password is wrong")
	}

	accessToken, err := helper.CreateToken(findUserIdentifier.Id, 1)

	if err != nil {
		panic(err)
	}

	expIn := 5

	result := response.LoginUserRes{
		AccessToken: accessToken,
		Type:        "Bearer",
		Exp:         expIn,
	}

	return result
}
