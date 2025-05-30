package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/LexusEgorov/api-calculator/internal/app"
	"github.com/LexusEgorov/api-calculator/internal/config"
	"github.com/LexusEgorov/api-calculator/internal/logger"

	_ "github.com/LexusEgorov/api-calculator/docs"
)

// @title           API Calculator
// @version         1.0
// @description     Calculator which works through the API

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	cfg, err := config.New()

	if err != nil {
		log.Fatalf("Unable to load config: %v", err)
		return
	}

	log := logger.New(cfg.Env)
	app := app.New(log, cfg.Server.Port)

	app.Run()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)

	<-stopChan
	log.Info("Recieved interrupt signal")
	app.Stop()
}
