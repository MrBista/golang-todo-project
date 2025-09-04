package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MrBista/golang-todo-project/src/dto"
	"github.com/MrBista/golang-todo-project/src/services"
	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	GetUserByEmail(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
}

type userController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (s *userController) GetUserByEmail(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// w.Write([]byte("Get User By Email"))

	res := dto.ResponseHttp{
		Data:    "Get User By Email",
		Status:  http.StatusOK,
		Message: "Success",
	}

	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)

	encoder.Encode(res)

}
