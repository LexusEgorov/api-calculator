package main

import (
	"api-calculator/internal/app"
	"api-calculator/internal/logger"
)

func main() {
	log := logger.New()
	app := app.New(log)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
