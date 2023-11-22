package repository

import (
	"Qpay/internal/db"
	"Qpay/internal/model"
)

type PaymentGatewaysRepository interface {
	Create(paymentGateway model.PaymentGateways) (model.PaymentGateways, error)
	Update(paymentGateway model.PaymentGateways) (model.PaymentGateways, error)
	Delete(id uint) error
}

type paymentGatewaysRepository struct {
}

func NewPaymentGatewaysRepository() PaymentGatewaysRepository {
	return &paymentGatewaysRepository{}
}

func (u *paymentGatewaysRepository) Create(paymentGateway model.PaymentGateways) (model.PaymentGateways, error) {
	db, err := db.GetDatabaseConnection()

	if err != nil {
		return paymentGateway, err
	}

	tx := db.Create(&paymentGateway)
	return paymentGateway, tx.Error

}

func (u *paymentGatewaysRepository) Update(paymentGateway model.PaymentGateways) (model.PaymentGateways, error) {
	db, err := db.GetDatabaseConnection()

	if err != nil {
		return paymentGateway, err
	}

	tx := db.Create(&paymentGateway)
	return paymentGateway, tx.Error
}

func (u *paymentGatewaysRepository) Delete(id uint) error {
	db, err := db.GetDatabaseConnection()

	if err != nil {
		return err
	}

	tx := db.Delete(&model.PaymentGateways{}, id)
	return tx.Error
}
