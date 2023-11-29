package handler

import (
	"Qpay/internal/model"
	"Qpay/internal/repository"
	"Qpay/pkg/bank"
	"net/http"
	"time"

	"github.com/sinabakh/go-zarinpal-checkout"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PaymentHandler interface {
	PaymentRequest() echo.HandlerFunc
	PaymentVerification() echo.HandlerFunc
	PaymentCallback() echo.HandlerFunc
	PaymentAction() echo.HandlerFunc
}

type paymentHandler struct {
	config                    bank.Config
	transactionRepository     repository.TransactionRepository
	paymentGatewaysRepository repository.PaymentGatewaysRepository
}

func NewPaymentHandler(cfg bank.Config, transactionRepository repository.TransactionRepository, paymentGatewaysRepository repository.PaymentGatewaysRepository) paymentHandler {
	return paymentHandler{
		config:                    cfg,
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
			Description string `json:"description" form:"description" validate:"required"`
			Email       string `json:"email" form:"email" validate:"required"`
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
			Status:      "pending",
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

		if transaction.Status == "successful" {
			transaction.Status = "verified"
			updatedTransaction, err := s.transactionRepository.Update(transaction.ID, transaction)
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}

			if updatedTransaction.ID == 0 {
				return c.String(http.StatusInternalServerError, "transaction not updated")
			}
			return c.JSON(http.StatusOK, map[string]interface{}{
				"status": "verified",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "unsuccesful",
		})
	}
}

func (s *paymentHandler) PaymentCallback() echo.HandlerFunc {
	return func(c echo.Context) error {
		//get Authority from query string
		authority := c.QueryParam("Authority")
		//get Status from query string
		status := c.QueryParam("Status")

		transaction, err := s.transactionRepository.GetByBankAuthority(authority)

		if err != nil {
			// return c.String(http.StatusInternalServerError, err.Error())
			return c.Redirect(http.StatusFound, transaction.CallbackURL+"?status=unsuccessful"+"?authority="+transaction.Authority)
		}

		// if transaction not exists
		if transaction.ID == 0 {
			// return c.String(http.StatusNotFound, "Transaction not found")
			return c.Redirect(http.StatusFound, transaction.CallbackURL+"?status=unsuccessful"+"?authority="+transaction.Authority)
		}

		if status == "OK" {
			zarinPay, err := zarinpal.NewZarinpal(s.config.MerchantID, false)
			if err != nil {
				// return c.String(http.StatusInternalServerError, err.Error())
				return c.Redirect(http.StatusFound, transaction.CallbackURL+"?status=unsuccessful"+"?authority="+transaction.Authority)
			}
			authority := transaction.BankAuthority // The authority of the payment
			amount := transaction.Amount           // The amount of payment in Tomans
			_, refID, statusCode, err := zarinPay.PaymentVerification(amount, authority)

			if err != nil {
				// return c.String(http.StatusInternalServerError, "خطایی رخ داد: "+err.Error())
				return c.Redirect(http.StatusFound, transaction.CallbackURL+"?status=unsuccessful"+"?authority="+transaction.Authority)
			}

			if (statusCode == 101) || (statusCode == 100) {
				transaction.Status = "successful"
				transaction.TransactionTime = time.Now()
				transaction.BankRefID = refID
				updatedTransaction, err := s.transactionRepository.Update(transaction.ID, transaction)
				if err != nil {
					// return c.String(http.StatusInternalServerError, "transaction not updated")
					return c.Redirect(http.StatusFound, transaction.CallbackURL+"?status=unsuccessful"+"?authority="+transaction.Authority)
				}

				if updatedTransaction.ID == 0 {
					// return c.String(http.StatusInternalServerError, "transaction not updated")
					return c.Redirect(http.StatusFound, transaction.CallbackURL+"?status=unsuccessful"+"?authority="+transaction.Authority)
				}

				return c.Redirect(http.StatusFound, transaction.CallbackURL+"?status=successful"+"?authority="+transaction.Authority)
			}

		}

		// if status is not OK
		transaction.Status = "unsuccessful"
		updatedTransaction, err := s.transactionRepository.Update(transaction.ID, transaction)
		if err != nil {
			// return c.String(http.StatusInternalServerError, err.Error())
			return c.Redirect(http.StatusFound, transaction.CallbackURL+"?status=unsuccessful"+"?authority="+transaction.Authority)
		}

		if updatedTransaction.ID == 0 {
			// return c.String(http.StatusInternalServerError, "transaction not updated")
			return c.Redirect(http.StatusFound, transaction.CallbackURL+"?status=unsuccessful"+"?authority="+transaction.Authority)
		}

		return c.Redirect(http.StatusFound, transaction.CallbackURL+"?status=unsuccessful"+"?authority="+transaction.Authority)
	}
}

func (s *paymentHandler) PaymentAction() echo.HandlerFunc {
	return func(c echo.Context) error {
		paymentAuthority := c.Param("Authority")

		transaction, err := s.transactionRepository.GetByAuthority(paymentAuthority)

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// if transaction not exists
		if transaction.ID == 0 {
			return c.String(http.StatusNotFound, "Transaction not found")
		}

		zarinPay, err := zarinpal.NewZarinpal(s.config.MerchantID, false)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		paymentURL, authority, statusCode, err := zarinPay.NewPaymentRequest(transaction.Amount, s.config.BaseURL, transaction.Description, transaction.Email, transaction.Phone)

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		if statusCode == 100 {
			// redirect user to zarinpal
			transaction.BankAuthority = authority
			updatedTransaction, err := s.transactionRepository.Update(transaction.ID, transaction)
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}

			if updatedTransaction.ID == 0 {
				return c.String(http.StatusInternalServerError, "transaction not updated")
			}
			return c.Redirect(http.StatusFound, paymentURL)
		}

		return c.String(http.StatusInternalServerError, "خطایی رخ داد: "+string(statusCode))
	}
}
