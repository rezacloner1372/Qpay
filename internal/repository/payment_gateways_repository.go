package repository

import (
	"Qpay/internal/db"
	"Qpay/internal/model"
)

type PaymentGatewaysRepository interface {
	Create(user model.PaymentGateways) (model.PaymentGateways, error)
	Update(user model.PaymentGateways) (model.PaymentGateways, error)
	Delete(id uint) error
}

type paymentGatewaysRepository struct {
}

func NewPaymentGatewaysRepository() PaymentGatewaysRepository {
	return &paymentGatewaysRepository{}
}

func (u *paymentGatewaysRepository) Create(user model.PaymentGateways) (model.PaymentGateways, error) {
	db := db.GetDatabaseConnection()
	err := db.Create(&user).Error
	return user, err
}

func (u *paymentGatewaysRepository) Update(user model.PaymentGateways) (model.PaymentGateways, error) {
	db := db.GetDatabaseConnection()
	err := db.Save(&user).Error
	return user, err
}

func (u *paymentGatewaysRepository) Delete(id uint) error {
	db := db.GetDatabaseConnection()
	err := db.Delete(&model.PaymentGateways{}, id).Error
	return err
}
