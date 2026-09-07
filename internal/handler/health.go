package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nalendro16/learn-vibe-code/internal/dto"
	"gorm.io/gorm"
)

type HealthHandler struct {
	db *gorm.DB
}

func NewHealthHandler(db *gorm.DB) *HealthHandler {
	return &HealthHandler{db: db}
}

func (h *HealthHandler) HealthCheck(c *gin.Context) {
	dbStatus := "connected"
	if h.db != nil {
		sqlDB, err := h.db.DB()
		if err != nil || sqlDB.Ping() != nil {
			dbStatus = "disconnected"
		}
	} else {
		dbStatus = "uninitialized"
	}

	data := gin.H{
		"database":  dbStatus,
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	}

	dto.SendOK(c, "Service is healthy", data)
}
