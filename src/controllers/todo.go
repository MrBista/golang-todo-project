package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MrBista/golang-todo-project/helper"
	"github.com/MrBista/golang-todo-project/src/dto/request"
	"github.com/MrBista/golang-todo-project/src/dto/response"
	"github.com/MrBista/golang-todo-project/src/services"
	"github.com/julienschmidt/httprouter"
)

type TodoController interface {
	CreateTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	UpdateTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	DeleteByIdTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	DeleteAllTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindByIdTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindAllTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params)
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

func (s *TodoControllerImpl) UpdateTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	updateBody := request.TodoReq{}
	todoIdParam := params.ByName("todoId")
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&updateBody)

	if err != nil {
		panic(err)
	}

	todoId, err := strconv.Atoi(todoIdParam)

	if err != nil {
		panic(err)
	}

	resTodo := s.TodoService.Update(r.Context(), updateBody, todoId)

	webResponse := response.CommonResponse{
		Data:    resTodo,
		Status:  http.StatusOK,
		Message: "Successfully update todo",
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)

	encoder.Encode(webResponse)

}

func (s *TodoControllerImpl) DeleteByIdTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (s *TodoControllerImpl) DeleteAllTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (s *TodoControllerImpl) FindByIdTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}

func (s *TodoControllerImpl) FindAllTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("not implemented") // TODO: Implement
}
