package repository

import (
	"Qpay/internal/db"
	"Qpay/internal/model"

	"errors"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction model.Transactions) (model.Transactions, error)
	Update(id uint, transaction model.Transactions) (model.Transactions, error)
	Delete(id uint) error
	GetAll() ([]model.Transactions, error)
	GetById(id uint) (model.Transactions, error)
}

type transactionRepository struct {
}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{}
}

func (t *transactionRepository) Create(transaction model.Transactions) (model.Transactions, error) {
	db, err := db.GetDatabaseConnection()
	if err != nil {
		return transaction, err
	}
	tx := db.Create(&transaction)
	return transaction, tx.Error
}

func (t *transactionRepository) Update(id uint, transaction model.Transactions) (model.Transactions, error) {

	db, err := db.GetDatabaseConnection()

	if err != nil {
		return transaction, err
	}
	tx := db.Model(&transaction).Where("id = ?", id).Updates(transaction)
	return transaction, tx.Error
}

func (t *transactionRepository) Delete(id uint) error {
	db, err := db.GetDatabaseConnection()

	if err != nil {
		return err
	}
	tx := db.Delete(&model.Transactions{}, id)
	return tx.Error
}

func (t *transactionRepository) GetAll() ([]model.Transactions, error) {

	var transactions []model.Transactions
	db, err := db.GetDatabaseConnection()

	if err != nil {
		return transactions, err
	}
	tx := db.Find(&transactions)
	return transactions, tx.Error
}

func (t *transactionRepository) GetById(id uint) (model.Transactions, error) {

	var transaction model.Transactions
	db, err := db.GetDatabaseConnection()

	if err != nil {
		return transaction, err
	}
	tx := db.First(&transaction, id)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return transaction, nil
	}
	return transaction, tx.Error
}
