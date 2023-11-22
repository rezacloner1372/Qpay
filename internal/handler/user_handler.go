package handler

import (
	"Qpay/internal/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	Signup() echo.HandlerFunc
	Login() echo.HandlerFunc
	Logout() echo.HandlerFunc
}

type userHandler struct {
	repository repository.UserRepository
}

func NewUserHandler(userRepository repository.UserRepository) userHandler {
	return userHandler{
		repository: userRepository,
	}
}

func (s *userHandler) Signup() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Signup")
	}
}

func (s *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Login")
	}
}

func (s *userHandler) Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Logout")
	}
}
