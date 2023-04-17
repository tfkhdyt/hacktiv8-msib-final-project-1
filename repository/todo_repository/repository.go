package todo_repository

import (
	"hacktiv8-msib-final-project-1/entity"
	"hacktiv8-msib-final-project-1/pkg/errs"
)

type TodoRepository interface {
	CreateTodo(todo *entity.Todo) (*entity.Todo, errs.MessageErr)
	GetAllTodos() ([]entity.Todo, errs.MessageErr)
	GetTodoByID(id uint) (*entity.Todo, errs.MessageErr)
	DeleteTodo(id uint) errs.MessageErr
}
