package repositiry

import (
	"github.com/Krabik6/todo-list"
	"github.com/jmoiron/sqlx"
)

type Authorisation interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorisation
	TodoList
	TodoItem
}

func NewRepositiry(db *sqlx.DB) *Repository {
	return &Repository{
		Authorisation: NewAuthPostgres(db),
	}
}
