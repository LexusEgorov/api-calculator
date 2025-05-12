package app

import (
	"github.com/sirupsen/logrus"

	"api-calculator/internal/calculator"
	"api-calculator/internal/server"
	"api-calculator/internal/storage/requests"
)

type App struct {
	s *server.Server
	l *logrus.Logger
}

func New(logger *logrus.Logger) *App {
	//TODO: cache
	reqStorage := requests.New()

	controller := calculator.New(logger, nil, reqStorage)
	s := server.New(controller)

	return &App{
		s: s,
		l: logger,
	}
}

func (a App) Run() error {
	a.l.Info("Starting app...")
	return a.s.Run()
}
