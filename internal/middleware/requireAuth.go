package middleware

import (
	"Qpay/internal/repository"

	"Qpay/pkg/jwt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type MassageResponse struct {
	Message string `json:"message"`
}

func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, MassageResponse{Message: "Error Token is empty"})
		}
		splitedToken := strings.Split(string(tokenString), " ")
		if len(splitedToken) < 2 {
			return c.JSON(http.StatusUnauthorized, MassageResponse{Message: "Unauthorized"})
		}
		tokenString = splitedToken[1]

		claims, err := jwt.ParseJWT(tokenString)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, MassageResponse{Message: "Unauthorized"})
		}

		userRepository := repository.NewUserRepository()
		user, err := userRepository.Find(uint(claims["id"].(float64)))

		if err != nil {
			return c.JSON(http.StatusUnauthorized, MassageResponse{Message: "Error while getting user"})
		}

		if user.ID == 0 {
			return c.JSON(http.StatusUnauthorized, MassageResponse{Message: "Cannot find user"})
		}

		c.Set("user", user)
		return next(c)
	}
}
