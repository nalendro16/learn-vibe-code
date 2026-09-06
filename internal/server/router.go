package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nalendro16/learn-vibe-code/internal/config"
	"github.com/nalendro16/learn-vibe-code/internal/handler"
	"gorm.io/gorm"
)

func NewRouter(cfg *config.Config, db *gorm.DB) *gin.Engine {
	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	healthHandler := handler.NewHealthHandler(db)

	// Base health check
	router.GET("/health", healthHandler.HealthCheck)

	// API routes group
	api := router.Group("/api/v1")
	{
		api.GET("/health", healthHandler.HealthCheck)
	}

	return router
}
