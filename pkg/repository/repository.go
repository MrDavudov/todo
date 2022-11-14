package repository

import (
	"github.com/MrDavudov/todo/internal/model"
	"github.com/jmoiron/sqlx"
)

type Auth interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
	FindUser(user, username string) error
}

type TodoList interface {
	Create(userId int, list model.TodoList) (int, error)
	GetAll(userId int) ([]model.TodoList, error)
	GetById(userId, listId int) (model.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input model.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item model.TodoItem) (int, error)
	GetAll(userId, listId int) ([]model.TodoItem, error)
	GetById(userId, itemId int) (model.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input model.UpdateItemInput) error
}

type Repository struct {
	Auth
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthPostgres(db),
		TodoList: NewTodoListPostgres(db),
		TodoItem: NewTodoItemPostgres(db),
	}
}
