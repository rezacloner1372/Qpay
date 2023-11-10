package model

import (
	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name      string `gorm:"unique_index;not null"`
	Family    string `gorm:"unique_index;not null"`
	Email     string `gorm:"unique_index;not null"`
	Cellphone string `gorm:"unique_index;not null"`
	Username  string `gorm:"unique_index;not null"`
	Password  string `gorm:"not null"`
	Status    string `gorm:"not null"`
	RoleId    int
	Type      *string
}

func (u *User) HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}

func (u *User) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
	return err == nil
}
