package app

import (
	"github.com/sirupsen/logrus"

	"api-calculator/internal/calculator"
	"api-calculator/internal/server"
)

type App struct {
	s *server.Server
	l *logrus.Logger
}

func New(logger *logrus.Logger) *App {
	//TODO: cache
	//TODO: storage

	controller := calculator.New(logger, nil, nil)
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
