package handler

import (
	"Qpay/internal/model"
	"Qpay/internal/repository"
	"Qpay/pkg/utils"
	"fmt"
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
		email := c.FormValue("email")
		password := c.FormValue("password")

		hashedPassword, err := utils.HashPassword(password)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Registration failed")
		}

		user := model.User{Email: email, Password: hashedPassword}
		result, err := s.repository.Create(user)
		fmt.Println(111, result, err)

		if err != nil {
			return c.String(http.StatusInternalServerError, "Registration failed")
		}

		return c.JSON(http.StatusOK, result)
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
