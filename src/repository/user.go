package repository

import (
	"context"
	"database/sql"

	"github.com/MrBista/golang-todo-project/src/model"
)

type UserRepositry interface {
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) model.User
	FindByEmailOrUsername(ctx context.Context, tx *sql.Tx, identifier string) (model.User, error)
	CreateUser(ctx context.Context, tx *sql.Tx, user model.User) model.User
}
