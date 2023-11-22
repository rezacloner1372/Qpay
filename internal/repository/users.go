package repository

import (
	"Qpay/internal/model"
	"Qpay/pkg/utils"
	"errors"
	"gorm.io/gorm"
)

func GetUserByEmail(db *gorm.DB, email string) (*model.User, error) {
	var dbUser model.User
	result := db.First(&dbUser, "email = ?", email)
	if result.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	return &dbUser, nil
}

func RegisterUser(db *gorm.DB, email string, password string) (*model.User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := model.User{Email: email, Password: hashedPassword}
	result := db.Create(&user)
	if err = result.Error; err != nil {
		return nil, err
	}

	return &user, nil
}
