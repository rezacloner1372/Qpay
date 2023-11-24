package repository

import (
	"Qpay/internal/db"
	"Qpay/internal/model"

	"errors"

	"gorm.io/gorm"
)

type TariffRepository interface {
	Create(tariff model.Tariffs) (model.Tariffs, error)
	Update(id uint, ariff model.Tariffs) (model.Tariffs, error)
	Delete(id uint) error
	GetAll() ([]model.Tariffs, error)
	GetById(id uint) (model.Tariffs, error)
}

type tariffRepository struct {
}

func NewTariffRepository() TariffRepository {
	return &tariffRepository{}
}

func (t *tariffRepository) Create(tariff model.Tariffs) (model.Tariffs, error) {
	db, err := db.GetDatabaseConnection()
	if err != nil {
		return tariff, err
	}
	tx := db.Create(&tariff)
	return tariff, tx.Error
}

func (t *tariffRepository) Update(id uint, tariff model.Tariffs) (model.Tariffs, error) {
	db, err := db.GetDatabaseConnection()

	if err != nil {
		return tariff, err
	}
	tx := db.Model(&tariff).Where("id = ?", id).Updates(tariff)
	return tariff, tx.Error
}

func (t *tariffRepository) Delete(id uint) error {
	db, err := db.GetDatabaseConnection()

	if err != nil {
		return err
	}
	tx := db.Delete(&model.Tariffs{}, id)
	return tx.Error
}

func (t *tariffRepository) GetAll() ([]model.Tariffs, error) {

	var tariffs []model.Tariffs
	db, err := db.GetDatabaseConnection()

	if err != nil {
		return tariffs, err
	}
	tx := db.Find(&tariffs)
	return tariffs, tx.Error
}

func (t *tariffRepository) GetById(id uint) (model.Tariffs, error) {

	var tariff model.Tariffs
	db, err := db.GetDatabaseConnection()

	if err != nil {
		return tariff, err
	}
	tx := db.First(&tariff, id)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return tariff, nil
	}
	return tariff, tx.Error
}
