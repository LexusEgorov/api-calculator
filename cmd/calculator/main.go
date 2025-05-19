package main

import (
	"log"

	"github.com/LexusEgorov/api-calculator/internal/app"
	"github.com/LexusEgorov/api-calculator/internal/config"
	"github.com/LexusEgorov/api-calculator/internal/logger"
)

//TODO: Add unit tests
//TODO: Move to Echo
//TODO: Add documentation

func main() {
	cfg, err := config.New()

	if err != nil {
		log.Fatalf("Unable to load config: %v", err)
		return
	}

	log := logger.New(cfg.Env)
	app := app.New(log, cfg.Server.Port)

	if err := app.Run(); err != nil {
		log.Fatalf("Can't start application: %v", err)
	}
}
