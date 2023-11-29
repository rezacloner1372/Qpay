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
	paymentGatewayHandler := handler.NewPaymentGatewaysHandler(paymentGatewaysRepo)
	e.POST("/payment-gateways/new", paymentGatewayHandler.Create(), customeMiddleware.RequireAuth)
	e.PUT("/payment-gateways/:id", paymentGatewayHandler.Update(), customeMiddleware.RequireAuth)
	e.DELETE("/payment-gateways/:id", paymentGatewayHandler.Delete(), customeMiddleware.RequireAuth)
	e.GET("/payment-gateways/all", paymentGatewayHandler.GetAll(), customeMiddleware.RequireAuth)
	e.GET("/payment-gateways/:id", paymentGatewayHandler.GetById(), customeMiddleware.RequireAuth)

	tariffRepo := repository.NewTariffRepository()
	tariffHandler := handler.NewTariffHandler(tariffRepo)
	e.POST("/tariff/new", tariffHandler.Create(), customeMiddleware.RequireAuth)
	e.PUT("/tariff/:id", tariffHandler.Update(), customeMiddleware.RequireAuth)
	e.DELETE("/tariff/:id", tariffHandler.Delete(), customeMiddleware.RequireAuth)
	e.GET("/tariff/all", tariffHandler.GetAll(), customeMiddleware.RequireAuth)
	e.GET("/tariff/:id", tariffHandler.GetById(), customeMiddleware.RequireAuth)

	transactionRepo := repository.NewTransactionRepository()
	transactionHandler := handler.NewTransactionHandler(transactionRepo)
	e.POST("/transaction/new", transactionHandler.Create(), customeMiddleware.RequireAuth)
	e.PUT("/transaction/:id", transactionHandler.Update(), customeMiddleware.RequireAuth)
	e.DELETE("/transaction/:id", transactionHandler.Delete(), customeMiddleware.RequireAuth)
	e.GET("/transaction/all", transactionHandler.GetAll(), customeMiddleware.RequireAuth)
	e.GET("/transaction/:id", transactionHandler.GetById(), customeMiddleware.RequireAuth)

	paymentHandler := handler.NewPaymentHandler(transactionRepo, paymentGatewaysRepo)
	e.POST("/payment/request", paymentHandler.PaymentRequest())
	e.POST("/payment/verify", paymentHandler.PaymentVerification())
	e.GET("/payment/callback", paymentHandler.PaymentCallback())
	e.GET("/payment/:Authority", paymentHandler.PaymentAction())
}
