package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/asylcreek/simple-auth-server/user"
)

type SignUpPostData struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required,eqfield=Password"`
}

func signUp(db *gorm.DB, context *gin.Context) {
	var data SignUpPostData

	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})

		return
	}

	isUsernameTaken := user.DoesUsernameExist(db, data.Username)

	if isUsernameTaken {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "This username has already been taken, please use a different one.",
		})

		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Something went really wrong",
		})

		return
	}

	newUser := user.User{Username: data.Username, Password: string(hashedPassword)}

	if err := user.AddUser(db, &newUser); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Something went really wrong",
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   newUser,
	})
}
