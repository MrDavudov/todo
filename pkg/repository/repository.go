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
}

type TodoItem interface {
}

type Repository struct {
	Auth
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthPostgres(db),
	}
}
