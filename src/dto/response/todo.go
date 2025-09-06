package response

import "time"

type TodoResponse struct {
	Id          int
	Title       string
	Description string
	Status      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
