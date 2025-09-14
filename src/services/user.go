package services

import (
	"context"
	"database/sql"

	"github.com/MrBista/golang-todo-project/helper"
	"github.com/MrBista/golang-todo-project/src/dto/request"
	"github.com/MrBista/golang-todo-project/src/dto/response"
	"github.com/MrBista/golang-todo-project/src/exception"
	"github.com/MrBista/golang-todo-project/src/model"
	"github.com/MrBista/golang-todo-project/src/repository"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	RegisterUser(ctx context.Context, req request.RegisterUserRequest) (response.RegisterUserResponse, error)
	LoginUser(ctx context.Context, req request.LoginUserReq) (response.LoginUserRes, error)
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

func (s *UserServiceImpl) RegisterUser(ctx context.Context, reqBody request.RegisterUserRequest) (response.RegisterUserResponse, error) {

	err := s.validateRegisterRequest(reqBody)
	if err != nil {
		return response.RegisterUserResponse{}, err
	}
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

	userSaved, err := s.userRepository.CreateUser(ctx, trx, user)

	if err != nil {
		return response.RegisterUserResponse{}, err
	}

	userResponse := response.RegisterUserResponse{
		Id:       userSaved.Id,
		Email:    userSaved.Email,
		Username: userSaved.Username,
		FullName: userSaved.FullName,
	}

	return userResponse, nil
}

func (s *UserServiceImpl) LoginUser(ctx context.Context, req request.LoginUserReq) (response.LoginUserRes, error) {

	// 1. cari dulu user nya
	// 2. kalau ga ada maka panic
	// 3. kala ada maka check dulu passwordnya
	// 4. kalau ga match maka panic
	// 5. genereate jwt dengan user id dan exp time dari viper

	if err := s.validateLogin(req); err != nil {
		return response.LoginUserRes{}, err
	}

	trx, err := s.DB.Begin()

	if err != nil {
		helper.Logger().Error(err)
		panic(err)
	}

	findUserIdentifier, err := s.userRepository.FindByEmailOrUsername(ctx, trx, req.Identifier)

	if err != nil {
		helper.Logger().Error(err)
		errMessage := err.Error()
		return response.LoginUserRes{}, exception.NewDbError(errMessage)
	}

	if err := helper.ComparePassword(req.Password, findUserIdentifier.Password); err != nil {

		return response.LoginUserRes{}, exception.NewBadReqeust("username/password is wrong")
	}

	accessToken, err := helper.CreateToken(findUserIdentifier.Id, 1)

	if err != nil {

		return response.LoginUserRes{}, exception.NewBadReqeust(err.Error())
	}

	expIn := 5

	result := response.LoginUserRes{
		AccessToken: accessToken,
		Type:        "Bearer",
		Exp:         expIn,
	}

	return result, nil
}

// validateRegisterRequest validates the register request
func (s *UserServiceImpl) validateRegisterRequest(req request.RegisterUserRequest) error {
	fieldErrors := make(map[string]string)

	// Example validations - adjust according to your struct
	if req.Email == "" {
		fieldErrors["email"] = "Email is required"
	}
	if req.Username == "" {
		fieldErrors["username"] = "Username is required"
	}
	if req.Password == "" {
		fieldErrors["password"] = "Password is required"
	} else if len(req.Password) < 6 {
		fieldErrors["password"] = "Password must be at least 6 characters"
	}

	if len(fieldErrors) > 0 {
		return exception.NewValidationError("Validation failed", fieldErrors)
	}

	return nil
}

func (s *UserServiceImpl) validateLogin(req request.LoginUserReq) error {
	fieldErrors := make(map[string]string)

	if req.Identifier == "" {
		fieldErrors["identifier"] = "Identifier is required"
	}

	if req.Password == "" {
		fieldErrors["password"] = "password is required"
	}

	if len(fieldErrors) > 0 {
		return exception.NewValidationError("Validation error", fieldErrors)
	}

	return nil
}
