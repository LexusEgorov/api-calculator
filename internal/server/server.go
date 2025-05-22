package echosrv

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	mdw "github.com/LexusEgorov/api-calculator/internal/middleware"
)

type CalcHandler interface {
	HandleHistory(ctx echo.Context) error
	HandleSum(ctx echo.Context) error
	HandleMult(ctx echo.Context) error
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

	server.POST("/sum", handler.HandleSum, middleware.WithLogging, middleware.WithAuth)
	server.POST("/mult", handler.HandleMult, middleware.WithLogging, middleware.WithAuth)
	server.GET("/history", handler.HandleHistory, middleware.WithLogging, middleware.WithAuth)

	return &Server{
		handler: handler,
		logger:  logger,
		server:  server,
	}
}

func (s Server) Run() {
	s.logger.Infof("Server is running on: %d port", s.port)
	if err := s.server.Start(fmt.Sprintf(":%d", s.port)); err != nil {
		s.logger.Fatalf("Server starting error: %v", err)
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
