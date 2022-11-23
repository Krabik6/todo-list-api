package service

import (
	"github.com/Krabik6/todo-list"
	"github.com/Krabik6/todo-list/pkg/repositiry"
)

type Authorisation interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorisation
	TodoList
	TodoItem
}

func NewSrvice(repos *repositiry.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(repos),
	}
}
