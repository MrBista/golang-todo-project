package repository

import (
	"context"
	"database/sql"

	"github.com/MrBista/golang-todo-project/helper"
	"github.com/MrBista/golang-todo-project/src/model"
)

type Todo interface {
	Create(ctx context.Context, sql *sql.Tx, todo *model.Todo)
	Update(ctx context.Context, sql *sql.Tx, todo model.Todo)
	FindById(ctx context.Context, sql *sql.Tx, id int)
	FindAll(ctx context.Context, sql *sql.Tx, userId int)
	DeleteById(ctx context.Context, sql *sql.Tx, id int)
	DeleteAll(ctx context.Context, sql *sql.Tx, userId int)
}

type TodoImpl struct {
}

func NewTodo() Todo {
	return &TodoImpl{}
}

func (r *TodoImpl) Create(ctx context.Context, sql *sql.Tx, todo *model.Todo) {
	SQL := "INSERT INTO todos(user_id, title, description, status, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)"

	res, err := sql.ExecContext(ctx, SQL, todo.UserId, todo.Title, todo.Description, todo.Status, todo.CreatedAt, todo.UpdatedAt)

	if err != nil {
		helper.Logger().Error(err)
		panic(err)
	}
	lastId, err := res.LastInsertId()

	if err != nil {
		helper.Logger().Error(err)
		panic(err)
	}

	todo.Id = int(lastId)
}
func (r *TodoImpl) Update(ctx context.Context, sql *sql.Tx, todo model.Todo) {

}

func (r *TodoImpl) FindById(ctx context.Context, sql *sql.Tx, id int) {

}

func (r *TodoImpl) FindAll(ctx context.Context, sql *sql.Tx, userId int) {

}

func (r *TodoImpl) DeleteAll(ctx context.Context, sql *sql.Tx, userId int) {

}

func (r *TodoImpl) DeleteById(ctx context.Context, sql *sql.Tx, id int) {

}
