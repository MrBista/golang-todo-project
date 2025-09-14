package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MrBista/golang-todo-project/helper"
	"github.com/MrBista/golang-todo-project/src/dto/request"
	"github.com/MrBista/golang-todo-project/src/dto/response"
	"github.com/MrBista/golang-todo-project/src/exception"
	"github.com/MrBista/golang-todo-project/src/handler"
	"github.com/MrBista/golang-todo-project/src/services"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

type UserController interface {
	GetUserByEmail(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	UserRegister(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	LoginUser(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type UserControllerImpl struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (s *UserControllerImpl) GetUserByEmail(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func (s *UserControllerImpl) UserRegister(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// 1. decode dulu req body nya
	// 2. call register
	// 3. encode response

	logger := logrus.New()
	reqUserBody := request.RegisterUserRequest{}

	decode := json.NewDecoder(r.Body)

	err := decode.Decode(&reqUserBody)
	if err != nil {
		logger.Error(err)
		handler.HandleError(w, exception.NewBadReqeust("Invalid JSON format body"))
		return
	}

	userRes, err := s.UserService.RegisterUser(r.Context(), reqUserBody)
	logger.Error(err)
	if err != nil {
		handler.HandleError(w, err)
		return
	}
	responseResult := response.CommonResponse{
		Data:    userRes,
		Status:  http.StatusOK,
		Message: "Successfully register user",
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	encoder := json.NewEncoder(w)

	err = encoder.Encode(responseResult)

	if err != nil {
		logrus.Error(err)
		panic(err)
	}

}

func (s *UserControllerImpl) LoginUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	reqLogin := request.LoginUserReq{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&reqLogin)

	loginUser, err := s.UserService.LoginUser(r.Context(), reqLogin)

	if err != nil {
		handler.HandleError(w, err)
		return
	}

	resResult := response.CommonResponse{
		Data:    loginUser,
		Status:  http.StatusOK,
		Message: "Successfully login",
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)

	err = encoder.Encode(resResult)

	if err != nil {
		helper.Logger().Error(err)
		panic(err)
	}

}
