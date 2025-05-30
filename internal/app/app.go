package app

import (
	"context"
	"time"

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

func (a App) Run() {
	a.logger.Info("Starting app...")
	go a.server.Run()
}

func (a App) Stop() {
	a.logger.Info("Stopping app...")

	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	doneCh := make(chan error)
	go func() {
		doneCh <- a.server.Stop(ctx)
	}()

	select {
	case err := <-doneCh:
		if err != nil {
			a.logger.Errorf("Error while stopping server: %v", err)
		}
		a.logger.Info("App has been stopped gracefully")

	case <-ctx.Done():
		a.logger.Warn("App stopped forced")
	}
}
