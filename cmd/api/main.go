package main

import (
	"context"

	"github.com/brandaoca44/infraforge/internal/config"
	"github.com/brandaoca44/infraforge/internal/database"
	"github.com/brandaoca44/infraforge/internal/server"
	"github.com/brandaoca44/infraforge/internal/worker"
)

func main() {
	cfg := config.Load()

	database.Connect()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if cfg.MonitorEnabled {
	    worker.StartMonitor(ctx, cfg)
}

	app := server.New(cfg)

	server.Run(cfg, app, cancel)
}