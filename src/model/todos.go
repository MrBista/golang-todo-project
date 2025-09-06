package model

import "time"

type Todo struct {
	Id          int
	UserId      int
	Title       string
	Description string
	Status      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
