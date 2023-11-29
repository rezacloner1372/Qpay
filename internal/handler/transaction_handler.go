package handler

import (
	"Qpay/internal/model"
	"Qpay/internal/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransactionHandler interface {
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetById() echo.HandlerFunc
}

type transactionHandler struct {
	repository repository.TransactionRepository
}

func NewTransactionHandler(transactionRepository repository.TransactionRepository) transactionHandler {
	return transactionHandler{
		repository: transactionRepository,
	}
}

func (s *transactionHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Parse request body to extract transaction information
		var newTransaction model.Transactions
		if err := c.Bind(&newTransaction); err != nil {
			return c.String(http.StatusBadRequest, "Invalid request body")
		}
		// Create transaction using transactionRepository
		createdTransaction, err := s.repository.Create(newTransaction)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, createdTransaction)
	}
}

func (s *transactionHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		transactionId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid request")
		}
		transactionIdUint := uint(transactionId)

		// Parse request body to extract transaction information
		var newTransaction model.Transactions
		if err := c.Bind(&newTransaction); err != nil {
			return c.String(http.StatusBadRequest, "Invalid request body")
		}

		// Update transaction using transactionRepository
		updatedTransaction, err := s.repository.Update(transactionIdUint, newTransaction)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, updatedTransaction)
	}
}

func (s *transactionHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		transactionId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid request")
		}
		transactionIdUint := uint(transactionId)

		// Delete transaction using transactionRepository
		err = s.repository.Delete(transactionIdUint)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.String(http.StatusOK, "Transaction deleted successfully")
	}
}

func (s *transactionHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get all transactions using transactionRepository
		transactions, err := s.repository.GetAll()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, transactions)
	}
}

func (s *transactionHandler) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		transactionId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid request")
		}
		transactionIdUint := uint(transactionId)

		// Get transaction by id using transactionRepository
		transaction, err := s.repository.GetById(transactionIdUint)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, transaction)
	}
}
