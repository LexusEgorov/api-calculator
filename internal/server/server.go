package server

import (
	"net/http"

	"api-calculator/internal/calculator"
)

type Server struct {
	c *calculator.CalcController
	s *http.Server
}

func New(controller *calculator.CalcController) *Server {
	server := Server{
		c: controller,
		s: nil, //TODO
	}

	http.HandleFunc("/sum", controller.HandleSum)
	http.HandleFunc("/mult", controller.HandleMult)
	http.HandleFunc("/story", controller.HandleHistory)

	return &server
}

func (s Server) Run() error {
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		return err
	}

	return nil
}
