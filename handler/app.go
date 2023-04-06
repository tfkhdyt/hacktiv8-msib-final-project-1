package handler

import (
	"hacktiv8-msib-final-project-1/database"
	"hacktiv8-msib-final-project-1/handler/http_handler"
	"hacktiv8-msib-final-project-1/repository/todo_repository/todo_pg"
	"hacktiv8-msib-final-project-1/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var PORT = os.Getenv("PORT")

func StartApp() {
	if PORT == "" {
		PORT = "8080"
	}
	r := gin.Default()

	db := database.GetPostgresInstance()
	todoRepo := todo_pg.NewTodoPG(db)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := http_handler.NewTodoHandler(todoService)

	r.POST("/todos", todoHandler.CreateTodo)

	log.Fatalln(r.Run(":" + PORT))
}
