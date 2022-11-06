package servise

import "github.com/MrDavudov/todo/pkg/repository"

type Auth interface {
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

func NewService(rep *repository.Repository) *Service {
	return &Service{}
}