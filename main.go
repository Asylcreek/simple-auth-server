package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/asylcreek/simple-auth-server/database"
)

func main() {
	envLoadErr := godotenv.Load()

	if envLoadErr != nil {
		panic("Failed to load environment variables")
	}

	database.ConnectDB()

	router := gin.Default()

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	if err := router.Run(":4500"); err != nil {
		panic("Something went wrong")
	}
}
