package server

import (
	"Qpay/internal/db"
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

	db.CreateDBConnection()
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

	paymentGatewaysRepo := repository.NewPaymentGatewaysRepository()
	paymentHandler := handler.NewPaymentGatewaysHandler(paymentGatewaysRepo)
	e.POST("/payment/gateway/new", paymentHandler.CreatePersonalGateway())
	e.POST("/payment/gateway/business/new", paymentHandler.CreateBusinessGateway())

	tariffRepo := repository.NewTariffRepository()
	tariffHandler := handler.NewTariffHandler(tariffRepo)
	e.POST("/tariff/new", tariffHandler.Create())
	e.PUT("/tariff/:id", tariffHandler.Update())
	e.DELETE("/tariff/:id", tariffHandler.Delete())
	e.GET("/tariff/all", tariffHandler.GetAll())
	e.GET("/tariff/:id", tariffHandler.GetById())
}
