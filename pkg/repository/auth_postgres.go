package repository

import (
	"fmt"

	"github.com/MrDavudov/todo/internal/model"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, nil
	}
	
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}

func (r *AuthPostgres) FindUser(name, username string) error {
	var user model.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE name=$1 AND username=$2", usersTable)
	err := r.db.Get(&user, query, name, username)
	fmt.Println(err)

	return err
}