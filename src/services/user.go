package services

import (
	"context"

	"github.com/MrBista/golang-todo-project/src/model"
	"github.com/MrBista/golang-todo-project/src/repository"
)

type UserService interface {
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}

type userService struct {
	userRepository repository.UserRepositry
}

func NewUserService(userRepo repository.UserRepositry) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (s *userService) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	s.userRepository.FindByEmail(ctx, email)

	return nil, nil
}
