package test

import (
	"Qpay/internal/handler"
	"Qpay/internal/model"
	"Qpay/internal/repository"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUserSignup(t *testing.T) {
	e := echo.New()

	userRepo := repository.NewUserRepository()
	userHandler := handler.NewUserHandler(userRepo)

	t.Run("Valid Request", func(t *testing.T) {
		requestPayload := `{
			"name": "John",
			"family": "Doe",
			"email": "john.doe@example.com",
			"cellphone": "123456789",
			"username": "johndoe",
			"password": "password123"
		}`

		req := httptest.NewRequest(http.MethodPost, "/auth/signup", bytes.NewBufferString(requestPayload))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := userHandler.Signup()(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Validation Error", func(t *testing.T) {
		requestPayload := `{
			"family": "Doe",
			"email": "john.doe@example.com",
			"cellphone": "123456789",
			"username": "johndoe",
			"password": "password123"
		}`

		req := httptest.NewRequest(http.MethodPost, "/auth/signup", bytes.NewBufferString(requestPayload))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := userHandler.Signup()(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("Duplicate Username", func(t *testing.T) {
		existingUser := model.User{
			Name:      "John",
			Family:    "Doe",
			Email:     "existing.user@example.com",
			Cellphone: "987654321",
			Username:  "johndoe",
			Password:  "hashedpassword",
		}
		userRepo.Create(existingUser)

		requestPayload := `{
			"name": "Jane",
			"family": "Doe",
			"email": "jane.doe@example.com",
			"cellphone": "123456789",
			"username": "johndoe",
			"password": "password123"
		}`

		req := httptest.NewRequest(http.MethodPost, "/auth/signup", bytes.NewBufferString(requestPayload))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := userHandler.Signup()(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

