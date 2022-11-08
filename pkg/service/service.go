package servise

import (
	"github.com/MrDavudov/todo/internal/model"
	"github.com/MrDavudov/todo/pkg/repository"
)

type Auth interface {
	CreateUser(user model.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Auth
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repo.Auth),
	}
}

