package app

import (
	"log"

	cfg "github.com/kangman53/project-sprint-halo-suster/config"
	"github.com/kangman53/project-sprint-halo-suster/database"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func StartApp() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		Prefork:     cfg.Prefork,
	})

	dbPool := database.GetConnPool()
	// Temporary helper to initiate tables
	if err := database.InitiateTables(dbPool); err != nil {
		log.Fatal("Error when initializing tables:", err)
	}
	defer dbPool.Close()
	app.Use(logger.New())
	// Register BP
	RegisterBluePrint(app, dbPool)

	err := app.Listen(":8080")
	log.Fatal(err)
}
