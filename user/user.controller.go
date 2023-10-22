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

func GetUserByUsername(db *gorm.DB, username string) (User, error) {
	var foundUser User

	result := db.Model(User{}).Where("username = ?", username).First(&foundUser)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return User{}, errors.New("We cannot seem to find that user")
	}

	return foundUser, nil
}
