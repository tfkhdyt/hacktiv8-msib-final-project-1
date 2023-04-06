package handler

import (
	"hacktiv8-msib-final-project-1/database"
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

	log.Fatalln(r.Run(":" + PORT))
}
