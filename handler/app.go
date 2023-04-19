package handler

import (
	"hacktiv8-msib-final-project-1/database"
	"hacktiv8-msib-final-project-1/docs"
	"hacktiv8-msib-final-project-1/handler/http_handler"
	"hacktiv8-msib-final-project-1/repository/todo_repository/todo_pg"
	"hacktiv8-msib-final-project-1/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var PORT = os.Getenv("PORT")

func init() {
	docs.SwaggerInfo.Title = "Todo Application"
	docs.SwaggerInfo.Description = "This is a todo list management application"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func StartApp() {
	db := database.GetPostgresInstance()

	if PORT == "" {
		PORT = "8080"
	}
	r := gin.Default()

	todoRepo := todo_pg.NewTodoPG(db)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := http_handler.NewTodoHandler(todoService)

	r.POST("/todos", todoHandler.CreateTodo)
	r.GET("/todos", todoHandler.GetAllTodos)
	r.GET("/todos/:id", todoHandler.GetTodoByID)
	r.PUT("/todos/:id", todoHandler.UpdateTodo)
	r.DELETE("/todos/:id", todoHandler.DeleteTodo)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatalln(r.Run(":" + PORT))
}
