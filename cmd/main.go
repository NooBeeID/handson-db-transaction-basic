package main

import (
	"handson-db-transactions/infra/database"
	"handson-db-transactions/internal/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	filename := "config.yaml"
	err := config.LoadConfigFromYaml(filename)
	if err != nil {
		panic(err)
	}

	cfg := config.GetConfig()
	router := fiber.New(fiber.Config{
		Prefork: cfg.App.Prefork,
		AppName: cfg.App.Name,
	})

	// init database
	db, err := database.ConnectPostgre(cfg.DB)
	if err != nil {
		panic(err)
	}

	_ = db

	if err := router.Listen(cfg.App.Port); err != nil {
		panic(err)
	}
}
