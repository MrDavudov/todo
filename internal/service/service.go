package service

import (
	"github.com/MrDavudov/todo/internal/model"
	"github.com/MrDavudov/todo/internal/repository"
)

type Auth interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list model.TodoList) (int, error)
	GetAll(userId int) ([]model.TodoList, error)
	GetById(userId, listId int) (model.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input model.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, itemId int, item model.TodoItem) (int, error)
	GetAll(userId, listId int) ([]model.TodoItem, error)
	GetById(userId, itemId int) (model.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input model.UpdateItemInput) error
}

type Service struct {
	Auth
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repo.Auth),
		TodoList: NewTodoListService(repo.TodoList),
		TodoItem: NewTodoItemService(repo.TodoItem, repo.TodoList),
	}
}

