package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MrBista/golang-todo-project/helper"
	"github.com/MrBista/golang-todo-project/src/dto/request"
	"github.com/MrBista/golang-todo-project/src/dto/response"
	"github.com/MrBista/golang-todo-project/src/services"
	"github.com/julienschmidt/httprouter"
)

type TodoController interface {
	CreateTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type TodoControllerImpl struct {
	TodoService services.Todo
}

func NewTodoController(todoService services.Todo) TodoController {
	return &TodoControllerImpl{
		TodoService: todoService,
	}
}

func (s *TodoControllerImpl) CreateTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todoReq := request.TodoReq{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&todoReq)
	if err != nil {
		helper.Logger().Error(err)
		panic(err)
	}

	result := s.TodoService.Create(r.Context(), todoReq)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)

	responseWeb := response.CommonResponse{
		Data:    result,
		Status:  http.StatusCreated,
		Message: "Successfully create todo",
	}

	encoder.Encode(responseWeb)

}
