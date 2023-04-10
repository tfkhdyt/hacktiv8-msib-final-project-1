package todo_pg

import (
	"hacktiv8-msib-final-project-1/entity"
	"hacktiv8-msib-final-project-1/pkg/errs"
	"hacktiv8-msib-final-project-1/repository/todo_repository"
	"log"

	"gorm.io/gorm"
)

type todoPG struct {
	db *gorm.DB
}

func NewTodoPG(db *gorm.DB) todo_repository.TodoRepository {
	return &todoPG{db: db}
}

func (t *todoPG) CreateTodo(todo *entity.Todo) (*entity.Todo, errs.MessageErr) {
	if err := t.db.Create(todo).Error; err != nil {
		log.Println(err.Error())
		return nil, errs.NewInternalServerError("Failed to create new todo")
	}

	return todo, nil
}

func (t *todoPG) GetAllTodos() ([]entity.Todo, errs.MessageErr) {
	var todos []entity.Todo

	if err := t.db.Find(&todos).Error; err != nil {
		log.Println(err.Error())
		return nil, errs.NewInternalServerError("Failed to get all todos")
	}

	return todos, nil
}
