package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/MrBista/golang-todo-project/src/model"
)

type UserRepositryImpl struct {
}

func NewUserRepository() UserRepositry {
	return &UserRepositryImpl{}
}

func (repo *UserRepositryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) model.User {

	script := "select id, username, email from users where email = ?"

	row, err := tx.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer row.Close()

	user := model.User{}
	if row.Next() {

		err := row.Scan(&user.Id, &user.Username, &user.Email)

		if err != nil {
			panic(err)
		}
	}

	return user
}

func (repo *UserRepositryImpl) FindByEmailOrUsername(ctx context.Context, tx *sql.Tx, identifier string) (model.User, error) {
	SQL := "SELECT id, username, email, password FROM users WHERE email = ? OR username = ?"

	row, err := tx.QueryContext(ctx, SQL, identifier, identifier)

	if err != nil {
		panic(err)
	}

	defer row.Close()

	user := model.User{}

	if row.Next() {
		err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user not found")
		} else if err != nil {
			return user, err
		}
	}

	return user, nil
}

func (repo *UserRepositryImpl) CreateUser(ctx context.Context, tx *sql.Tx, user model.User) model.User {
	SQL := "insert into users(username, email, password, full_name) values(?, ?, ?, ?)"

	res, err := tx.ExecContext(ctx, SQL, user.Username, user.Email, user.Password, user.FullName)

	if err != nil {
		panic(err)
	}

	lastIdInserted, err := res.LastInsertId()

	if err != nil {
		panic(err)
	}
	user.Id = int(lastIdInserted)
	return user
}
