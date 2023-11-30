package handler

import (
	"Qpay/internal/model"
	"Qpay/internal/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TariffHandler interface {
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetById() echo.HandlerFunc
}

type tariffHandler struct {
	repository repository.TariffRepository
}

func NewTariffHandler(tariffRepository repository.TariffRepository) tariffHandler {
	return tariffHandler{
		repository: tariffRepository,
	}
}

func (s *tariffHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Parse request body to extract tariff information
		// var newTariff model.Tariffs
		var request struct {
			Name         string `json:"name" form:"name" validate:"required"`
			Description  string `json:"description" form:"description" validate:"required"`
			Price        int    `json:"price" form:"price" validate:"required"`
			Currency     string `json:"currency" form:"currency" validate:"required"`
			ValidityDays int    `json:"validity_days" form:"validity_days" validate:"required"`
			IsDefault    int    `json:"is_default" form:"is_default" validate:"required"`
		}
		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid body request")
		}

		if err := c.Validate(request); err != nil {
			return c.JSON(http.StatusBadRequest, "Validation error: "+err.Error())
		}

		newTariff := model.Tariffs{
			Name:         request.Name,
			Description:  request.Description,
			Price:        request.Price,
			Currency:     request.Currency,
			ValidityDays: request.ValidityDays,
			IsDefault:    request.IsDefault,
		}

		// Create tariff using tariffRepository
		createdTariff, err := s.repository.Create(newTariff)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, createdTariff)
	}
}

func (s *tariffHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		tariffId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid request")
		}
		tariffIdUint := uint(tariffId)

		var request struct {
			Name         string `json:"name" form:"name" validate:"required"`
			Description  string `json:"description" form:"description" validate:"required"`
			Price        int    `json:"price" form:"price" validate:"required"`
			Currency     string `json:"currency" form:"currency" validate:"required"`
			ValidityDays int    `json:"validity_days" form:"validity_days" validate:"required"`
			IsDefault    int    `json:"is_default" form:"is_default" validate:"required"`
		}
		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid body request")
		}

		if err := c.Validate(request); err != nil {
			return c.JSON(http.StatusBadRequest, "Validation error: "+err.Error())
		}

		newTariff := model.Tariffs{
			Name:         request.Name,
			Description:  request.Description,
			Price:        request.Price,
			Currency:     request.Currency,
			ValidityDays: request.ValidityDays,
			IsDefault:    request.IsDefault,
		}

		// Update tariff using tariffRepository
		updatedTariff, err := s.repository.Update(tariffIdUint, newTariff)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error updating tariff")
		}

		return c.JSON(http.StatusOK, updatedTariff)
	}
}

func (s *tariffHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Parse request body to extract tariff information
		tariffId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid request")
		}
		tariffIdUint := uint(tariffId)

		// Delete tariff using tariffRepository
		delErr := s.repository.Delete(tariffIdUint)
		if delErr != nil {
			return c.String(http.StatusInternalServerError, "Error deleting tariff")
		}

		return c.JSON(http.StatusOK, "Tariff deleted")
	}
}

func (s *tariffHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get all tariffs using tariffRepository
		tariffs, err := s.repository.GetAll()
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error getting all tariffs")
		}

		return c.JSON(http.StatusOK, tariffs)
	}
}

func (s *tariffHandler) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Parse request body to extract tariff information
		tariffId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid request")
		}
		tariffIdUint := uint(tariffId)

		// Get tariff by id using tariffRepository
		tariff, err := s.repository.GetById(tariffIdUint)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error getting tariff by id")
		}

		return c.JSON(http.StatusOK, tariff)
	}
}
