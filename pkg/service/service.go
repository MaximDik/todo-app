package service

import "github.com/MaximDik/todo-app/pkg/repository"

type Authorization interface {
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

func NewService(rep *repository.Repository) *Service {
	return &Service{}
}
