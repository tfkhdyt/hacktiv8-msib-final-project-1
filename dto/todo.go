package dto

import "hacktiv8-msib-final-project-1/entity"

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
