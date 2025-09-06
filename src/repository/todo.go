package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/MrBista/golang-todo-project/helper"
	"github.com/MrBista/golang-todo-project/src/model"
)

type Todo interface {
	Create(ctx context.Context, sql *sql.Tx, todo *model.Todo) error
	Update(ctx context.Context, sql *sql.Tx, todo model.Todo) error
	FindById(ctx context.Context, sql *sql.Tx, id int) (model.Todo, error)
	FindAll(ctx context.Context, sql *sql.Tx, userId int) ([]model.Todo, error)
	DeleteById(ctx context.Context, sql *sql.Tx, id int) error
	DeleteAll(ctx context.Context, sql *sql.Tx, userId int) error
}

type TodoImpl struct {
}

func NewTodo() Todo {
	return &TodoImpl{}
}

func (r *TodoImpl) Create(ctx context.Context, sql *sql.Tx, todo *model.Todo) error {
	SQL := "INSERT INTO todos(user_id, title, description, status, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)"

	res, err := sql.ExecContext(ctx, SQL, todo.UserId, todo.Title, todo.Description, todo.Status, todo.CreatedAt, todo.UpdatedAt)

	if err != nil {
		helper.Logger().Error(err)
		return fmt.Errorf("terjadi kesalahan %w", err)
	}
	lastId, err := res.LastInsertId()

	if err != nil {
		helper.Logger().Error(err)
		return fmt.Errorf("terjadi kesalahan %w", err)
	}

	todo.Id = int(lastId)

	return nil
}
func (r *TodoImpl) Update(ctx context.Context, tx *sql.Tx, todo model.Todo) error {
	SQL := `UPDATE todos 
			SET title = ?,
				description = ?,
				status = ?,
				updated_at = ?
			WHERE id = ?`

	_, err := tx.ExecContext(ctx, SQL, todo.Title, todo.Description, todo.Status, time.Now(), todo.Id)

	if err != nil {
		helper.Logger().Error(err)
		return fmt.Errorf("terjadi kesalahan %w", err)
	}

	return nil

}

func (r *TodoImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (model.Todo, error) {
	SQL := `SELECT id, title, description, status, updated_at, created_at 
			FROM todos 
			WHERE id = ?`

	rows, err := tx.QueryContext(ctx, SQL, id)

	todo := model.Todo{}

	if err != nil {
		helper.Logger().Error(err)
		return todo, fmt.Errorf("terjadi kesalahan %w", err)

	}

	defer rows.Close()

	if !rows.Next() {
		return todo, fmt.Errorf("terjadi kesalahan %w", err)
	}

	err = rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Status, &todo.UpdatedAt, &todo.CreatedAt)
	if err == sql.ErrNoRows {
		helper.Logger().Error(err)
		return todo, fmt.Errorf("terjadi kesalahan %w", err)
	} else if err != nil {
		helper.Logger().Error(err)
		return todo, fmt.Errorf("terjadi kesalahan %w", err)
	}

	return todo, nil
}

func (r *TodoImpl) FindAll(ctx context.Context, tx *sql.Tx, userId int) ([]model.Todo, error) {
	SQL := `SELECT id, title, description, status, updated_at, created_at 
			FROM todos
			WHERE user_id = ?`

	rows, err := tx.QueryContext(ctx, SQL, userId)

	if err != nil {
		return nil, fmt.Errorf("terjadi kesalahan %w", err)
	}

	defer rows.Close()

	todos := []model.Todo{}

	for rows.Next() {
		todo := model.Todo{}

		err = rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Status, &todo.UpdatedAt, &todo.CreatedAt)
		if err == sql.ErrNoRows {
			helper.Logger().Error(err)
			return nil, fmt.Errorf("terjadi kesalahan %w", err)
		} else if err != nil {
			helper.Logger().Error(err)
			return nil, fmt.Errorf("terjadi kesalahan %w", err)
		}

		todos = append(todos, todo)

	}

	return todos, nil
}

func (r *TodoImpl) DeleteAll(ctx context.Context, tx *sql.Tx, userId int) error {
	SQL := `DELETE FROM todos where user_id = ?`

	_, err := tx.ExecContext(ctx, SQL, userId)

	if err != nil {
		return fmt.Errorf("terjadi kesalahan %w", err)
	}

	return nil
}

func (r *TodoImpl) DeleteById(ctx context.Context, tx *sql.Tx, id int) error {

	SQL := `DELETE FROM todos where user_id = ?`

	_, err := tx.ExecContext(ctx, SQL, id)

	if err != nil {
		return fmt.Errorf("terjadi kesalahan %w", err)
	}

	return nil
}
