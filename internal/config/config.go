package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port            string
	Env             string
	MonitorInterval time.Duration
	MonitorEnabled  bool
}

func Load() Config {
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	// intervalo
	intervalSec := 10
	if v := os.Getenv("MONITOR_INTERVAL_SECONDS"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			intervalSec = n
		}
	}

	// enable/disable
	monitorEnabled := true
	if v := os.Getenv("MONITOR_ENABLED"); v == "false" {
		monitorEnabled = false
	}

	return Config{
		Port:            port,
		Env:             env,
		MonitorInterval: time.Duration(intervalSec) * time.Second,
		MonitorEnabled:  monitorEnabled,
	}
}