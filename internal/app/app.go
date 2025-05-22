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

	timeDeadline := time.Second * 5
	deadline := time.Now().Add(timeDeadline)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)

	defer cancel()
	go func() {
		if err := a.server.Stop(ctx); err != nil {
			a.logger.Errorf("Error while stopping server: %v", err)
		}
	}()

	select {
	case <-ctx.Done():
		if ctx.Err() == context.DeadlineExceeded {
			a.logger.Warn("App has been stopped by deadline")
		} else {
			a.logger.Info("App has been stopped gracefully")
		}

	case <-time.After(timeDeadline):
		a.logger.Warn("App stopped forced")
	}
}
