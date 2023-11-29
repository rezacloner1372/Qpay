package handler

import (
	"Qpay/internal/model"
	"Qpay/internal/repository"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PaymentHandler interface {
	PaymentRequest() echo.HandlerFunc
	PaymentVerification() echo.HandlerFunc
	// PaymentCallback() echo.HandlerFunc
	PaymentAction() echo.HandlerFunc
}

type paymentHandler struct {
	transactionRepository     repository.TransactionRepository
	paymentGatewaysRepository repository.PaymentGatewaysRepository
}

func NewPaymentHandler(transactionRepository repository.TransactionRepository, paymentGatewaysRepository repository.PaymentGatewaysRepository) paymentHandler {
	return paymentHandler{
		transactionRepository:     transactionRepository,
		paymentGatewaysRepository: paymentGatewaysRepository,
	}
}

func (s *paymentHandler) PaymentRequest() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request struct {
			MerchantID  string `json:"merchant_id" form:"merchant_id" validate:"required"`
			Amount      int    `json:"amount" form:"amount" validate:"required"`
			CallbackURL string `json:"callback_url" form:"callback_url" validate:"required"`
			Description string `json:"description" form:"description"`
			Email       string `json:"email" form:"email"`
			Phone       string `json:"phone" form:"phone" validate:"required"`
		}

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid body request")
		}

		if err := c.Validate(request); err != nil {
			return c.JSON(http.StatusBadRequest, "Validation error: "+err.Error())
		}

		paymentGateway, err := s.paymentGatewaysRepository.GetByMerchantId(request.MerchantID)

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// if payment gateway not exists
		if paymentGateway.ID == 0 {
			return c.String(http.StatusNotFound, "Payment gateway not found")
		}

		transaction := model.Transactions{
			Amount:      request.Amount,
			CallbackURL: request.CallbackURL,
			Description: request.Description,
			Email:       request.Email,
			Phone:       request.Phone,
			Authority:   uuid.New().String(),
			GatewayID:   paymentGateway.ID,
			UserID:      paymentGateway.UserID,
		}

		createdTransaction, err := s.transactionRepository.Create(transaction)

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		
		return c.JSON(http.StatusOK, map[string]interface{}{
			"authority": createdTransaction.Authority,
		})
	}
}

func (s *paymentHandler) PaymentVerification() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request struct {
			MerchantID string `json:"merchant_id" form:"merchant_id" validate:"required"`
			Amount     int    `json:"amount" form:"amount" validate:"required"`
			Authority  string `json:"authority" form:"authority" validate:"required"`
		}

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid body request")
		}

		if err := c.Validate(request); err != nil {
			return c.JSON(http.StatusBadRequest, "Validation error: "+err.Error())
		}

		transaction, err := s.transactionRepository.GetByAuthority(request.Authority)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// if transaction not exists
		if transaction.ID == 0 {
			return c.String(http.StatusNotFound, "Transaction not found")
		}

		// if transaction exists but amount is not equal
		if transaction.Amount != request.Amount {
			return c.String(http.StatusBadRequest, "Invalid amount")
		}

		// if transaction exists but merchant_id is not equal
		if transaction.Gateway.MerchantID != request.MerchantID {
			return c.String(http.StatusBadRequest, "Invalid merchant_id")
		}

		return c.JSON(http.StatusOK, transaction)
	}
}

// func (s *paymentHandler) PaymentCallback() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 	}
// }

func (s *paymentHandler) PaymentAction() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Payment action")
	}
}
