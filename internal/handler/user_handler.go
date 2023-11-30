package handler

import (
	"Qpay/internal/model"
	"Qpay/internal/repository"
	"Qpay/pkg/jwt"
	"Qpay/pkg/utils"

	"net/http"
	"strings"

	"github.com/go-sql-driver/mysql"
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
		var request struct {
			Name      string `json:"name" form:"name" validate:"required"`
			Family    string `json:"family" form:"family" validate:"required"`
			Email     string `json:"email" form:"email" validate:"required,email"`
			Cellphone string `json:"cellphone" form:"cellphone" validate:"required,min=11,max=11,numeric"`
			Username  string `json:"username" form:"username" validate:"required"`
			Password  string `json:"password" form:"password" validate:"required,min=6"`
		}

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid body request")
		}

		if err := c.Validate(request); err != nil {
			return c.JSON(http.StatusBadRequest, "Validation error: "+err.Error())
		}

		hashedPassword, err := utils.HashPassword(request.Password)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Registration failed")
		}

		user := model.User{
			Name:      request.Name,
			Family:    request.Family,
			Email:     request.Email,
			Cellphone: request.Cellphone,
			Username:  request.Username,
			Password:  hashedPassword,
		}
		result, err := s.repository.Create(user)
		if err != nil {
			sqlError, ok := err.(*mysql.MySQLError)
			if ok && sqlError.Number == 1062 {
				if strings.Contains(sqlError.Message, "users.username") {
					return c.String(http.StatusBadRequest, "Username already exists")
				} else if strings.Contains(sqlError.Message, "users.email") {
					return c.String(http.StatusBadRequest, "Email already exists")
				} else if strings.Contains(sqlError.Message, "users.cellphone") {
					return c.String(http.StatusBadRequest, "Cellphone already exists")
				}
			}
			return c.String(http.StatusInternalServerError, "Error creating user")
		}

		return c.JSON(http.StatusOK, result)
	}
}

func (s *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request struct {
			Username string `json:"username" form:"username" validate:"required"`
			Password string `json:"password" form:"password" validate:"required"`
		}

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid body request")
		}

		if err := c.Validate(request); err != nil {
			return c.JSON(http.StatusBadRequest, "Validation error: "+err.Error())
		}
		user, err := s.repository.FindByUsername(request.Username)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error getting user")
		}

		if !user.CheckPassword(request.Password) {
			return c.String(http.StatusUnauthorized, "Invalid password")
		}

		token, err := jwt.GenerateJWT(user.ID)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error generating token")
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": token,
		})
	}
}

func (s *userHandler) Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Logout")
	}
}
