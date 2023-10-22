package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/asylcreek/simple-auth-server/user"
)

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

func login(db *gorm.DB, context *gin.Context) {
	var data LoginPostData

	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})

		return
	}

	user, err := user.GetUserByUsername(db, data.Username)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})

		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid password. Please try again.",
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}
