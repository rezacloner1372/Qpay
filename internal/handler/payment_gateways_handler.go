package handler

import (
	"Qpay/internal/model"
	"Qpay/internal/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PaymentGatewaysHandler interface {
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetById() echo.HandlerFunc
}

type paymentGatewaysHandler struct {
	repository repository.PaymentGatewaysRepository
}

func NewPaymentGatewaysHandler(paymentGatewaysRepository repository.PaymentGatewaysRepository) paymentGatewaysHandler {
	return paymentGatewaysHandler{
		repository: paymentGatewaysRepository,
	}
}

func (s *paymentGatewaysHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Parse request body to extract paymentGateways information
		var newPaymentGateways model.PaymentGateways
		if err := c.Bind(&newPaymentGateways); err != nil {
			return c.String(http.StatusBadRequest, "Invalid request body")
		}
		// Create paymentGateways using paymentGatewaysRepository
		createdPaymentGateways, err := s.repository.Create(newPaymentGateways)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, createdPaymentGateways)
	}
}

func (s *paymentGatewaysHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		paymentGatewaysId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid request")
		}
		paymentGatewaysIdUint := uint(paymentGatewaysId)

		// Parse request body to extract paymentGateways information
		var newPaymentGateways model.PaymentGateways
		if err := c.Bind(&newPaymentGateways); err != nil {
			return c.String(http.StatusBadRequest, "Invalid request body")
		}
		// Update paymentGateways using paymentGatewaysRepository
		updatedPaymentGateways, err := s.repository.Update(paymentGatewaysIdUint, newPaymentGateways)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, updatedPaymentGateways)
	}
}

func (s *paymentGatewaysHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		paymentGatewaysId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid request")
		}
		paymentGatewaysIdUint := uint(paymentGatewaysId)

		// Delete paymentGateways using paymentGatewaysRepository
		err = s.repository.Delete(paymentGatewaysIdUint)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusOK)
	}
}

func (s *paymentGatewaysHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get all paymentGateways using paymentGatewaysRepository
		paymentGateways, err := s.repository.GetAll()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, paymentGateways)
	}
}

func (s *paymentGatewaysHandler) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		paymentGatewaysId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid request")
		}
		paymentGatewaysIdUint := uint(paymentGatewaysId)

		// Get paymentGateways using paymentGatewaysRepository
		paymentGateways, err := s.repository.GetById(paymentGatewaysIdUint)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, paymentGateways)
	}
}
