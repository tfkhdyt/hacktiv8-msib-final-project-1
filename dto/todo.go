package dto

import (
	"hacktiv8-msib-final-project-1/entity"
	"time"
)

type NewTodoRequest struct {
	Title     string `json:"title" binding:"required" example:"Belajar Golang"`
	Completed bool   `json:"completed" example:"false"`
	UserID    uint   `json:"userId" binding:"required" example:"1"`
}

func (t *NewTodoRequest) TodoRequestToEntity() *entity.Todo {
	return &entity.Todo{
		Title:     t.Title,
		Completed: t.Completed,
		UserID:    t.UserID,
	}
}

type NewTodoResponse struct {
	Message string         `json:"message" example:"Todo with id 69 has been successfully created"`
	Data    NewTodoRequest `json:"data"`
}

type GetAllTodosResponse struct {
	Message string     `json:"message" example:"success"`
	Data    []TodoData `json:"data"`
}

type GetTodoByIDResponse struct {
	Message string           `json:"message" example:"success"`
	Data    TodoDataDetailed `json:"data"`
}

type TodoData struct {
	ID        uint   `json:"id" example:"69"`
	Title     string `json:"title" example:"Ngoding"`
	Completed bool   `json:"completed" example:"true"`
	UserID    uint   `json:"userId" example:"2"`
}

type TodoDataDetailed struct {
	ID        uint      `json:"id" example:"69"`
	Title     string    `json:"title" example:"Ngoding"`
	Completed bool      `json:"completed" example:"true"`
	UserID    uint      `json:"userId" example:"2"`
	CreatedAt time.Time `json:"createdAt" example:"2023-04-06T17:55:34.070213+07:00"`
	UpdatedAt time.Time `json:"updatedAt" example:"2023-04-06T17:55:34.070213+07:00"`
}
