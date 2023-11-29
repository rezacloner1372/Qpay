package test

import (
	"Qpay/internal/handler"
	"Qpay/internal/middleware"
	"Qpay/internal/model"
	"Qpay/internal/repository"
	"bytes"
	"encoding/json"
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

func TestUserLogin(t *testing.T) {
	e := echo.New()

	userRepo := repository.NewUserRepository()
	userHandler := handler.NewUserHandler(userRepo)

	testUser := model.User{
		Username: "testuser",
		Password: "testpassword",
	}
	userRepo.Create(testUser)

	t.Run("Valid Login", func(t *testing.T) {
		requestPayload := `{
			"username": "testuser",
			"password": "testpassword"
		}`

		req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBufferString(requestPayload))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := userHandler.Login()(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Invalid Password", func(t *testing.T) {
		requestPayload := `{
			"username": "testuser",
			"password": "wrongpassword"
		}`

		req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBufferString(requestPayload))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := userHandler.Login()(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	})

	t.Run("User Not Found", func(t *testing.T) {
		requestPayload := `{
			"username": "nonexistentuser",
			"password": "somepassword"
		}`

		req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBufferString(requestPayload))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := userHandler.Login()(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestUserLogout(t *testing.T) {
	e := echo.New()

	userRepo := repository.NewUserRepository()
	userHandler := handler.NewUserHandler(userRepo)

	e.Use(middleware.RequireAuth)

	e.GET("/auth/logout", userHandler.Logout())

	t.Run("Valid Logout with Authentication", func(t *testing.T) {
		loginRequestPayload := `{
			"username": "testuser",
			"password": "testpassword"
		}`

		loginReq := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBufferString(loginRequestPayload))
		loginReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		loginRec := httptest.NewRecorder()
		loginCtx := e.NewContext(loginReq, loginRec)

		err := userHandler.Login()(loginCtx)
		assert.NoError(t, err)

		var loginResponse map[string]string
		assert.NoError(t, json.Unmarshal(loginRec.Body.Bytes(), &loginResponse))
		authToken := loginResponse["token"]

		logoutReq := httptest.NewRequest(http.MethodGet, "/auth/logout", nil)
		logoutReq.Header.Set(echo.HeaderAuthorization, "Bearer "+authToken)
		logoutRec := httptest.NewRecorder()
		logoutCtx := e.NewContext(logoutReq, logoutRec)

		err = userHandler.Logout()(logoutCtx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, logoutRec.Code)
		assert.Equal(t, "Logout", logoutRec.Body.String())
	})

	t.Run("Unauthorized Logout without Authentication", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/auth/logout", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := userHandler.Logout()(c)

		assert.Error(t, err)
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	})
}
