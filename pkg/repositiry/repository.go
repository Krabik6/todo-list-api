package repositiry

import "github.com/jmoiron/sqlx"

type Authorisation interface {
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
	return &Repository{}
}
