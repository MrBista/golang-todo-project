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
	Update(ctx context.Context, todoReq request.TodoReq, todoId int) response.TodoResponse
	FindAll(ctx context.Context, userId int) []response.TodoResponse
	FindById(ctx context.Context, id int) response.TodoResponse
	DeleteAll(ctx context.Context, userId int)
	DeleteById(ctx context.Context, id int)
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

func (s *TodoImpl) Update(ctx context.Context, todoReq request.TodoReq, todoId int) response.TodoResponse {
	tx, err := s.DB.Begin()

	if err != nil {
		helper.Logger().Error(err)
		panic(err)
	}

	defer func() {
		err := recover()
		if err != nil {
			errRollback := tx.Rollback()
			if errRollback != nil {
				panic(errRollback)
			}
		} else {
			errCommit := tx.Commit()
			if errCommit != nil {
				panic(errCommit)
			}
		}
	}()

	s.FindById(ctx, todoId)

	todo := &model.Todo{
		Id:          todoId,
		Title:       todoReq.Title,
		Description: todoReq.Description,
		Status:      todoReq.Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	errTodo := s.TodoRepository.Update(ctx, tx, *todo)

	if errTodo != nil {
		helper.Logger().Error(errTodo)
		panic(errTodo)

	}

	updatedTodo := response.TodoResponse{
		Id:          todoId,
		Title:       todo.Title,
		Description: todo.Description,
		Status:      todo.Status,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}
	return updatedTodo

}

func (s *TodoImpl) FindAll(ctx context.Context, userId int) []response.TodoResponse {
	panic("not implemented") // TODO: Implement
}

func (s *TodoImpl) FindById(ctx context.Context, id int) response.TodoResponse {
	tx, err := s.DB.Begin()

	if err != nil {
		panic(err)
	}
	defer func() {
		err := recover()
		if err != nil {
			errRollback := tx.Rollback()
			if errRollback != nil {
				panic(errRollback)
			}
		} else {
			errCommit := tx.Commit()
			if errCommit != nil {
				panic(errCommit)
			}
		}
	}()
	res, err := s.TodoRepository.FindById(ctx, tx, id)

	if err != nil {
		helper.Logger().Error(err)
		panic(err)
	}
	return response.TodoResponse{
		Id:          res.Id,
		Title:       res.Title,
		Description: res.Description,
		Status:      res.Status,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}
}

func (s *TodoImpl) DeleteAll(ctx context.Context, userId int) {
	tx, err := s.DB.Begin()

	if err != nil {
		panic(err)
	}
	defer func() {
		err := recover()
		if err != nil {
			errRollback := tx.Rollback()
			if errRollback != nil {
				panic(errRollback)
			}
		} else {
			errCommit := tx.Commit()
			if errCommit != nil {
				panic(errCommit)
			}
		}
	}()

	err = s.TodoRepository.DeleteAll(ctx, tx, userId)
	if err != nil {
		panic(err)
	}
}

func (s *TodoImpl) DeleteById(ctx context.Context, id int) {
	tx, err := s.DB.Begin()

	if err != nil {
		panic(err)
	}
	defer func() {
		err := recover()
		if err != nil {
			errRollback := tx.Rollback()
			if errRollback != nil {
				panic(errRollback)
			}
		} else {
			errCommit := tx.Commit()
			if errCommit != nil {
				panic(errCommit)
			}
		}
	}()

	s.FindById(ctx, id)

	err = s.TodoRepository.DeleteById(ctx, tx, id)
	if err != nil {
		panic(err)
	}
}
