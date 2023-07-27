package database

import (
	"errors"
	"go_server/api/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUsers(db *gorm.DB) error {
	users := []models.User{
		{
			Email:    "email@email.com",
			Password: "123456",
			Name:     stringPointer("test"),
		},
		{
			Email:    "email2@email.com",
			Password: "123456",
			Name:     stringPointer("test2"),
		},
	}

	for _, user := range users {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)

		var existingUser models.User
		err = db.Where("email = ?", user.Email).First(&existingUser).Error
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}

			err = db.Create(&user).Error
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func stringPointer(s string) *string {
	return &s
}
