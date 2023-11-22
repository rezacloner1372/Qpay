package repository

import (
	"Qpay/internal/db"
	"Qpay/internal/model"
)

type UserRepository interface {
	Create(user model.User) (model.User, error)
	Update(user model.User) (model.User, error)
	Delete(id uint) error
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (u *userRepository) Create(user model.User) (model.User, error) {
	db := db.GetDatabaseConnection()
	err := db.Create(&user).Error
	return user, err
}

func (u *userRepository) Update(user model.User) (model.User, error) {
	db := db.GetDatabaseConnection()
	err := db.Save(&user).Error
	return user, err
}

func (u *userRepository) Delete(id uint) error {
	db := db.GetDatabaseConnection()
	err := db.Delete(&model.User{}, id).Error
	return err
}
