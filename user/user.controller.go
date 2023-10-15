package user

import (
	"errors"

	"gorm.io/gorm"
)

func DoesUsernameExist(db *gorm.DB, username string) bool {
	var foundUser User

	result := db.Model(User{}).Where("username = ?", username).First(&foundUser)

	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func AddUser(db *gorm.DB, user *User) error {
	result := db.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
