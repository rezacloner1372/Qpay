package repository

import (
	"Qpay/internal/model"
	"errors"

	"github.com/jinzhu/gorm"
)

func GetUserByEmail(db *gorm.DB, email string) (*model.User, error) {
	var dbUser model.User
	result := db.First(&dbUser, "email = ?", email)
	if result.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	return &dbUser, nil
}
