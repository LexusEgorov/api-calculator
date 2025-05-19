package app

import (
	"github.com/sirupsen/logrus"

	"github.com/LexusEgorov/api-calculator/internal/calculator"
	srv "github.com/LexusEgorov/api-calculator/internal/server"
	"github.com/LexusEgorov/api-calculator/internal/storage/cache"
	"github.com/LexusEgorov/api-calculator/internal/storage/requests"
)

type App struct {
	server *srv.Server
	logger *logrus.Logger
}

func New(logger *logrus.Logger, port int) *App {
	cacheStorage := cache.New()
	reqStorage := requests.New()

	handler := calculator.New(logger, cacheStorage, reqStorage)
	server := srv.New(handler, logger, port)

	return &App{
		server: server,
		logger: logger,
	}
}

func (a App) Run() error {
	a.logger.Info("Starting app...")
	return a.server.Run()
}
