package app

import (
	"github.com/sirupsen/logrus"

	"api-calculator/internal/calculator"
	srv "api-calculator/internal/server"
	"api-calculator/internal/storage/cache"
	"api-calculator/internal/storage/requests"
)

type App struct {
	server *srv.Server
	logger *logrus.Logger
}

func New(logger *logrus.Logger) *App {
	cacheStorage := cache.New()
	reqStorage := requests.New()

	handler := calculator.New(logger, cacheStorage, reqStorage)
	server := srv.New(handler, logger)

	return &App{
		server: server,
		logger: logger,
	}
}

func (a App) Run() error {
	a.logger.Info("Starting app...")
	return a.server.Run()
}
