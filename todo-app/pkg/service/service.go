package service

import (
	"github.com/owlsspell/todo-app"
	"github.com/owlsspell/todo-app/pkg/repository"
)

type Authorization interface {
	createUser(user todo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
