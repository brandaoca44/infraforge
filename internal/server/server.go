package server

import (
	"github.com/brandaoca44/infraforge/internal/config"
	"github.com/brandaoca44/infraforge/internal/health"
	"github.com/brandaoca44/infraforge/internal/services"

	"github.com/gin-gonic/gin"
)

func New(cfg config.Config) *gin.Engine {
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.GET("/health", health.Check)
	r.POST("/services", services.Create)
	r.GET("/services", services.List)

	return r
}