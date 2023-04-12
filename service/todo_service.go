package service

import (
	"fmt"
	"hacktiv8-msib-final-project-1/dto"
	"hacktiv8-msib-final-project-1/pkg/errs"
	"hacktiv8-msib-final-project-1/repository/todo_repository"
)

type TodoService interface {
	CreateTodo(payload *dto.NewTodoRequest) (*dto.NewTodoResponse, errs.MessageErr)
	GetAllTodos() (*dto.GetAllTodosResponse, errs.MessageErr)
	GetTodoByID(id uint) (*dto.GetTodoByIDResponse, errs.MessageErr)
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

func (t *todoService) GetAllTodos() (*dto.GetAllTodosResponse, errs.MessageErr) {
	todos, err := t.todoRepo.GetAllTodos()
	if err != nil {
		return nil, err
	}

	todoData := []dto.TodoData{}
	for _, todo := range todos {
		todoData = append(todoData, dto.TodoData{
			ID:        todo.ID,
			Title:     todo.Title,
			Completed: todo.Completed,
			UserID:    todo.UserID,
		})
	}

	response := &dto.GetAllTodosResponse{
		Message: "success",
		Data:    todoData,
	}

	return response, nil
}

func (t *todoService) GetTodoByID(id uint) (*dto.GetTodoByIDResponse, errs.MessageErr) {
	todo, err := t.todoRepo.GetTodoByID(id)
	if err != nil {
		return nil, err
	}

	response := &dto.GetTodoByIDResponse{
		Message: "success",
		Data: dto.TodoDataDetailed{
			ID:        todo.ID,
			Title:     todo.Title,
			Completed: todo.Completed,
			UserID:    todo.UserID,
			CreatedAt: todo.CreatedAt,
			UpdatedAt: todo.UpdatedAt,
		},
	}

	return response, nil
}
