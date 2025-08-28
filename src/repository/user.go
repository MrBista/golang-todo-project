package repository

import (
	"context"
	"database/sql"

	"github.com/MrBista/golang-todo-project/config"
	"github.com/MrBista/golang-todo-project/src/model"
)

type UserRepositry interface {
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}

type userRepositry struct {
	db *sql.DB
}

func NewUserRepository(db *config.Database) UserRepositry {
	return &userRepositry{
		db: db.DB,
	}
}

func (repo userRepositry) FindByEmail(ctx context.Context, email string) (*model.User, error) {

	script := "select id, username, email from users where email = ?"

	row, err := repo.db.QueryContext(ctx, script)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	user := model.User{}
	if row.Next() {

		err := row.Scan(&user.Id, &user.Username, &user.Email)

		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}
