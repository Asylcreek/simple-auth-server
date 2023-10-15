package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/asylcreek/simple-auth-server/auth"
	"github.com/asylcreek/simple-auth-server/database"
	"github.com/asylcreek/simple-auth-server/user"
)

func main() {
	envLoadErr := godotenv.Load()

	if envLoadErr != nil {
		panic("Failed to load environment variables")
	}

	db := database.Connect()

	err := db.AutoMigrate(&user.User{})
	if err != nil {
		panic(err.Error())
	}

	router := gin.Default()

	auth.Router(db, router)

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	if err := router.Run(":" + os.Getenv("PORT")); err != nil {
		panic("Something went wrong")
	}
}
