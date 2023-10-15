package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(db *gorm.DB, router *gin.Engine) {
	authRouter := router.Group("/auth")

	authRouter.POST("/signup", func(context *gin.Context) { signUp(db, context) })
}
