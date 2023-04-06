package service

import (
	"fmt"
	"hacktiv8-msib-final-project-1/dto"
	"hacktiv8-msib-final-project-1/pkg/errs"
	"hacktiv8-msib-final-project-1/repository/todo_repository"
)

type TodoService interface {
	CreateTodo(payload *dto.NewTodoRequest) (*dto.NewTodoResponse, errs.MessageErr)
}

type todoService struct {
	todoRepo todo_repository.TodoRepository
}

func NewTodoService(todoRepo todo_repository.TodoRepository) TodoService {
	return &todoService{todoRepo: todoRepo}
}

func (t *todoService) CreateTodo(payload *dto.NewTodoRequest) (*dto.NewTodoResponse, errs.MessageErr) {
	todo := payload.TodoRequestToEntity()

	createdTodo, err := t.todoRepo.CreateTodo(todo)
	if err != nil {
		return nil, err
	}

	response := &dto.NewTodoResponse{
		Message: fmt.Sprintf("Todo with id %v has been created successfully", createdTodo.ID),
		Data: dto.NewTodoRequest{
			Title:     createdTodo.Title,
			Completed: createdTodo.Completed,
			UserID:    createdTodo.UserID,
		},
	}

	return response, nil
}
