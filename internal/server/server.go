package server

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	mdw "github.com/LexusEgorov/api-calculator/internal/middleware"
)

type CalcHandler interface {
	HandleHistory(w http.ResponseWriter, r *http.Request)
	HandleSum(w http.ResponseWriter, r *http.Request)
	HandleMult(w http.ResponseWriter, r *http.Request)
}

type Server struct {
	handler CalcHandler
	logger  *logrus.Logger
	port    int
}

func New(handler CalcHandler, logger *logrus.Logger, port int) *Server {
	middleware := mdw.New(logger)
	server := Server{
		handler: handler,
		logger:  logger,
		port:    port,
	}

	http.HandleFunc("/sum", middleware.WithRecover(middleware.WithLogging(middleware.WithAuth(handler.HandleSum))))
	http.HandleFunc("/mult", middleware.WithRecover(middleware.WithLogging(middleware.WithAuth(handler.HandleMult))))
	http.HandleFunc("/history", middleware.WithRecover(middleware.WithLogging(middleware.WithAuth(handler.HandleHistory))))

	return &server
}

func (s Server) Run() error {
	s.logger.Infof("Server is running on localhost:%d", s.port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)

	if err != nil {
		return err
	}

	return nil
}
