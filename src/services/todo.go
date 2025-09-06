package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/MrBista/golang-todo-project/helper"
	"github.com/MrBista/golang-todo-project/src/dto/request"
	"github.com/MrBista/golang-todo-project/src/dto/response"
	"github.com/MrBista/golang-todo-project/src/model"
	"github.com/MrBista/golang-todo-project/src/repository"
)

type Todo interface {
	Create(ctx context.Context, todoReq request.TodoReq) response.TodoResponse
}

type TodoImpl struct {
	TodoRepository repository.Todo
	DB             *sql.DB
}

func NewTodo(TodoRepository repository.Todo, Db *sql.DB) Todo {

	return &TodoImpl{
		DB:             Db,
		TodoRepository: TodoRepository,
	}
}

func (s *TodoImpl) Create(ctx context.Context, todoRoq request.TodoReq) response.TodoResponse {

	trx, err := s.DB.Begin()

	if err != nil {
		helper.Logger().Error(err)
		panic(err)
	}

	defer func() {
		err := recover()
		if err != nil {
			errRolback := trx.Rollback()
			if errRolback != nil {
				panic(errRolback)
			}
		} else {
			errCommit := trx.Commit()
			if errCommit != nil {
				panic(errCommit)
			}
		}
	}()

	todo := &model.Todo{
		UserId:      1,
		Title:       todoRoq.Title,
		Description: todoRoq.Description,
		Status:      todoRoq.Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	// create todo
	s.TodoRepository.Create(ctx, trx, todo)

	todoRes := response.TodoResponse{
		Id:          todo.Id,
		Title:       todo.Title,
		Description: todo.Description,
		Status:      todo.Status,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}

	return todoRes
}
