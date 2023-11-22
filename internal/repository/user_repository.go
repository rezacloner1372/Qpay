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
	db, err := db.GetDatabaseConnection()

	if err != nil {
		return user, err
	}

	tx := db.Create(&user)
	return user, tx.Error
}

func (u *userRepository) Update(user model.User) (model.User, error) {
	db, err := db.GetDatabaseConnection()

	if err != nil {
		return user, err
	}

	tx := db.Create(&user)
	return user, tx.Error
}

func (u *userRepository) Delete(id uint) error {
	db, err := db.GetDatabaseConnection()

	if err != nil {
		return err
	}

	tx := db.Delete(&model.User{}, id)
	return tx.Error
}
