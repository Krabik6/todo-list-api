package service

import "github.com/Krabik6/todo-list/pkg/repositiry"

type Authorisation interface {
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
	return &Service{}
}
