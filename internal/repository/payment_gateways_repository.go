package repository

import (
	"Qpay/internal/db"
	"Qpay/internal/model"
)

type PaymentGatewaysRepository interface {
	Create(paymentGateway model.PaymentGateways) (model.PaymentGateways, error)
	Update(id uint, paymentGateway model.PaymentGateways) (model.PaymentGateways, error)
	Delete(id uint) error
	GetAll() ([]model.PaymentGateways, error)
	GetById(id uint) (model.PaymentGateways, error)
	GetByMerchantId(merchantId string) (model.PaymentGateways, error)
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

func (u *paymentGatewaysRepository) Update(id uint, paymentGateway model.PaymentGateways) (model.PaymentGateways, error) {
	db, err := db.GetDatabaseConnection()

	if err != nil {
		return paymentGateway, err
	}

	tx := db.Model(&paymentGateway).Where("id = ?", id).Updates(paymentGateway)
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

func (u *paymentGatewaysRepository) GetAll() ([]model.PaymentGateways, error) {

	var paymentGateways []model.PaymentGateways
	db, err := db.GetDatabaseConnection()

	if err != nil {
		return paymentGateways, err
	}

	tx := db.Preload("User").Preload("Tariff").Find(&paymentGateways)
	return paymentGateways, tx.Error
}

func (u *paymentGatewaysRepository) GetById(id uint) (model.PaymentGateways, error) {

	var paymentGateway model.PaymentGateways
	db, err := db.GetDatabaseConnection()

	if err != nil {
		return paymentGateway, err
	}

	tx := db.Preload("User").Preload("Tariff").Find(&paymentGateway, id)
	return paymentGateway, tx.Error
}

func (u *paymentGatewaysRepository) GetByMerchantId(merchantId string) (model.PaymentGateways, error) {

	var paymentGateway model.PaymentGateways
	db, err := db.GetDatabaseConnection()

	if err != nil {
		return paymentGateway, err
	}

	tx := db.Preload("User").Preload("Tariff").Find(&paymentGateway, "merchant_id = ?", merchantId)
	return paymentGateway, tx.Error
}
