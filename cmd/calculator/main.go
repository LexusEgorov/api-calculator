package main

import (
	"api-calculator/internal/app"
	"api-calculator/internal/logger"
	"os"
)

//TODO: Add unit tests
//TODO: Move to Echo
//TODO: Add documentation
//TODO: Add makefile for create documentation
//TODO: Add more logs
//TODO: Add middlewares: auth (check Id); logging (requests + codes)

func main() {
	deployment := os.Getenv("DEPLOYMENT")
	log := logger.New(deployment)
	app := app.New(log)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
