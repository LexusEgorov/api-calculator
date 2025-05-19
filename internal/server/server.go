package server

import (
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
	server  *http.Server
	logger  *logrus.Logger
}

func New(handler CalcHandler, logger *logrus.Logger) *Server {
	middleware := mdw.New(logger)
	server := Server{
		handler: handler,
		logger:  logger,
		server:  nil, //TODO
	}

	http.HandleFunc("/sum", middleware.WithLogging(middleware.WithAuth(handler.HandleSum)))
	http.HandleFunc("/mult", middleware.WithLogging(middleware.WithAuth(handler.HandleMult)))
	http.HandleFunc("/history", middleware.WithLogging(middleware.WithAuth(handler.HandleHistory)))

	return &server
}

func (s Server) Run() error {
	s.logger.Info("Server is running on localhost:8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		return err
	}

	return nil
}
