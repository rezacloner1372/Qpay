package server

import (
	"Qpay/internal/db"
	"Qpay/internal/handler"
	"Qpay/internal/repository"
	"net/http"

	customeMiddleware "Qpay/internal/middleware"

	"github.com/go-playground/validator"

	_ "Qpay/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Server struct {
	E *echo.Echo
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewServer() *Server {
	return &Server{
		E: echo.New(),
	}
}

func (s *Server) Start(address string, cfg *db.Config) error {
	e := s.E
	e.Validator = &CustomValidator{validator: validator.New()}

	// Middleware
	e.Use(middleware.Logger())  // Logger
	e.Use(middleware.Recover()) // Recover

	db.CreateDBConnection(cfg)
	routing(s.E)

	go func() {
		if err := e.Start(address); err != nil {
			e.Logger.Fatal(err)
		}
	}()

	return nil
}

func routing(e *echo.Echo) {
	userRepo := repository.NewUserRepository()
	userHandler := handler.NewUserHandler(userRepo)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/auth/signup", userHandler.Signup())
	e.POST("/auth/login", userHandler.Login())
	e.GET("/auth/logout", userHandler.Logout(), customeMiddleware.RequireAuth)

	paymentGatewaysRepo := repository.NewPaymentGatewaysRepository()
	paymentHandler := handler.NewPaymentGatewaysHandler(paymentGatewaysRepo)
	e.POST("/payment/gateway/new", paymentHandler.CreatePersonalGateway())
	e.POST("/payment/gateway/business/new", paymentHandler.CreateBusinessGateway())

	tariffRepo := repository.NewTariffRepository()
	tariffHandler := handler.NewTariffHandler(tariffRepo)
	e.POST("/tariff/new", tariffHandler.Create(), customeMiddleware.RequireAuth)
	e.PUT("/tariff/:id", tariffHandler.Update(), customeMiddleware.RequireAuth)
	e.DELETE("/tariff/:id", tariffHandler.Delete(), customeMiddleware.RequireAuth)
	e.GET("/tariff/all", tariffHandler.GetAll(), customeMiddleware.RequireAuth)
	e.GET("/tariff/:id", tariffHandler.GetById(), customeMiddleware.RequireAuth)
}
