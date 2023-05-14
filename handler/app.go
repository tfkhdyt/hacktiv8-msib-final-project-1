package handler

import (
	"fmt"
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

var (
	port    = os.Getenv("PORT")
	appHost = os.Getenv("APP_HOST")
)

func init() {
	if port == "" {
		port = "8080"
	}
	docs.SwaggerInfo.Title = "Todo Application"
	docs.SwaggerInfo.Description = "This is a todo list management application"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", appHost, port)
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func StartApp() {
	db := database.GetPostgresInstance()

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

	log.Fatalln(r.Run(":" + port))
}
