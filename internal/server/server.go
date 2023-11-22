package server

import (
	"Qpay/internal/handler"
	"Qpay/internal/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	E *echo.Echo
}

func NewServer() *Server {
	return &Server{
		E: echo.New(),
	}
}

func (s *Server) Start(address string) error {
	e := s.E

	// Middleware
	e.Use(middleware.Logger())  // Logger
	e.Use(middleware.Recover()) // Recover

	routing(s.E)

	if err := e.Start(address); err != nil {
		e.Logger.Fatal(err)
		return err
	}

	return nil
}

func routing(e *echo.Echo) {
	userRepo := repository.NewUserRepository()
	userHandler := handler.NewUserHandler(userRepo)

	e.POST("/auth/signup", userHandler.Signup())
	e.POST("/auth/login", userHandler.Login())
	e.GET("/auth/logout", userHandler.Logout())
}
