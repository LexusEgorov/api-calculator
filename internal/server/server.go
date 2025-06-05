package echosrv

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"

	mdw "github.com/LexusEgorov/api-calculator/internal/middleware"
)

type CalcHandler interface {
	HandleHistory(ctx echo.Context) error
	HandleSum(ctx echo.Context) error
	HandleMult(ctx echo.Context) error
	HandleCalculate(ctx echo.Context) error
}

type Server struct {
	handler CalcHandler
	server  *echo.Echo
	logger  *logrus.Logger
	port    int
}

func New(handler CalcHandler, logger *logrus.Logger, port int) *Server {
	middleware := mdw.New(logger)
	server := echo.New()

	server.Use(middleware.WithRecover, middleware.WithLogging, middleware.WithAuth)

	server.POST("/sum", handler.HandleSum)
	server.POST("/mult", handler.HandleMult)
	server.POST("/calc", handler.HandleCalculate)
	server.GET("/history", handler.HandleHistory)

	server.GET("/swagger/*", echoSwagger.WrapHandler)

	return &Server{
		handler: handler,
		logger:  logger,
		server:  server,
		port:    port,
	}
}

func (s Server) Run() {
	s.logger.Infof("Server is running on: localhost:%d", s.port)
	if err := s.server.Start(fmt.Sprintf("localhost:%d", s.port)); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			s.logger.Fatalf("Server starting error: %v", err)
		}
	}
}

func (s Server) Stop(ctx context.Context) error {
	s.logger.Info("Stopping server...")
	err := s.server.Shutdown(ctx)

	if err != nil {
		s.logger.Error(err)
	}

	return err
}
