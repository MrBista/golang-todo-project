package request

type RegisterUserRequest struct {
	Username string `json:"username" validate:"required, min=3, max=30"`
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required, min=1"`
	FullName string `json:"fullName"`
}
