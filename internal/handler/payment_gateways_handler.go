package handler

import (
	"Qpay/internal/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PaymentGatewaysHandler interface {
	CreatePersonalGateway() echo.HandlerFunc
	CreateBusinessGateway() echo.HandlerFunc
}

type paymentGatewaysHandler struct {
	repository repository.PaymentGatewaysRepository
}

func NewPaymentGatewaysHandler(paymentGatewaysRepository repository.PaymentGatewaysRepository) paymentGatewaysHandler {
	return paymentGatewaysHandler{
		repository: paymentGatewaysRepository,
	}
}

func (s *paymentGatewaysHandler) CreatePersonalGateway() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "CreatePersonalGateway")
	}
}

func (s *paymentGatewaysHandler) CreateBusinessGateway() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "CreateBusinessGateway")
	}
}
