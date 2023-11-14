package server

import (
	"fmt"

	"github.com/labstack/echo"
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
	fmt.Println(address)
	if err := s.E.Start(address); err != nil {
		s.E.Logger.Fatal(err)
		return err
	}

	return nil
}

func routing(e *echo.Echo) {
}
