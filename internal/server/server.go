package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	r.DELETE("/services/:id", services.Delete)

	return r
}

func Run(cfg config.Config, r *gin.Engine, stopBackgroundJobs context.CancelFunc) {
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	go func() {
		log.Printf("InfraForge API running on port %s", cfg.Port)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	stopBackgroundJobs()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("forced shutdown: %v", err)
	}

	log.Println("Server stopped cleanly.")
}