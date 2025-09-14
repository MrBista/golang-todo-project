package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MrBista/golang-todo-project/src/dto/request"
	"github.com/MrBista/golang-todo-project/src/dto/response"
	"github.com/MrBista/golang-todo-project/src/exception"
	"github.com/MrBista/golang-todo-project/src/handler"
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
		errRes := exception.NewBadReqeust(fmt.Errorf("terjadi kesalahan %w", err).Error())
		handler.HandleError(w, errRes)
		return
	}

	result, err := s.TodoService.Create(r.Context(), todoReq)

	if err != nil {
		handler.HandleError(w, err)
		return
	}

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
		errRes := exception.NewBadReqeust(fmt.Errorf("terjadi kesalahan %w", err).Error())
		handler.HandleError(w, errRes)
		return
	}

	todoId, err := strconv.Atoi(todoIdParam)

	if err != nil {
		panic(err)
	}

	resTodo, err := s.TodoService.Update(r.Context(), updateBody, todoId)

	if err != nil {
		handler.HandleError(w, err)
		return
	}

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
	todoIdParam := params.ByName("todoId")

	todoId, err := strconv.Atoi(todoIdParam)

	if err != nil {
		panic(err)
	}

	s.TodoService.DeleteById(r.Context(), todoId)

	webResponse := response.CommonResponse{
		Data:    true,
		Message: "Successfully delete todo",
		Status:  http.StatusOK,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	encoder.Encode(webResponse)
}

func (s *TodoControllerImpl) DeleteAllTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	// todo ambil user id dari context middleware
	userId := 1
	s.TodoService.DeleteById(r.Context(), userId)

	webResponse := response.CommonResponse{
		Data:    true,
		Message: "Successfully delete todo",
		Status:  http.StatusOK,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	encoder.Encode(webResponse)
}

func (s *TodoControllerImpl) FindByIdTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todoIdParam := params.ByName("todoId")

	todoId, err := strconv.Atoi(todoIdParam)

	if err != nil {
		panic(err)
	}

	result := s.TodoService.FindById(r.Context(), todoId)

	webResponse := response.CommonResponse{
		Data:    result,
		Message: "Successfully find todo",
		Status:  http.StatusOK,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	encoder.Encode(webResponse)
}

func (s *TodoControllerImpl) FindAllTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// todo ambil user id dari context middleware
	userId := 1

	result := s.TodoService.FindAll(r.Context(), userId)

	webResponse := response.CommonResponse{
		Data:    result,
		Message: "Successfully find todo",
		Status:  http.StatusOK,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	encoder.Encode(webResponse)
}
