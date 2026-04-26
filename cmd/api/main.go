package main

import (
	"log"

	"github.com/brandaoca44/infraforge/internal/config"
	"github.com/brandaoca44/infraforge/internal/server"
)

func main() {
	cfg := config.Load()

	app := server.New(cfg)

	log.Printf("InfraForge API running on port %s", cfg.Port)

	if err := app.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}