package model

import "time"

type User struct {
	Id       int
	Username string
	Email    string
	Password string
	FullName string
	CreateAt time.Time
	UpdateAt time.Time
}
