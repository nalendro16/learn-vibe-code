package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nalendro16/learn-vibe-code/internal/config"
	"github.com/nalendro16/learn-vibe-code/internal/handler"
	"github.com/nalendro16/learn-vibe-code/internal/repository"
	"github.com/nalendro16/learn-vibe-code/internal/service"
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

	// Initialize Handlers
	healthHandler := handler.NewHealthHandler(db)
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Base routes
	router.GET("/health", healthHandler.HealthCheck)

	// API routes
	api := router.Group("/api")
	{
		api.GET("/health", healthHandler.HealthCheck)
		api.POST("/register", userHandler.Register)

		v1 := api.Group("/v1")
		{
			v1.GET("/health", healthHandler.HealthCheck)
			v1.POST("/register", userHandler.Register)
		}
	}

	return router
}
