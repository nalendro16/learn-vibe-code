package handler

import (
	"log/slog"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/nalendro16/learn-vibe-code/internal/config"
	"github.com/nalendro16/learn-vibe-code/internal/database"
	"github.com/nalendro16/learn-vibe-code/internal/server"
	"gorm.io/gorm"
)

var (
	appRouter *gin.Engine
	appDB     *gorm.DB
	initOnce  sync.Once
)

func initApp() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("Failed to load configuration", "error", err)
	}

	appDB, err = database.NewPostgresDB(cfg.DB)
	if err != nil {
		slog.Warn("Could not connect to PostgreSQL database", "error", err)
	}

	appRouter = server.NewRouter(cfg, appDB)
}

// Handler is the entrypoint for Vercel Serverless Function
func Handler(w http.ResponseWriter, r *http.Request) {
	initOnce.Do(initApp)
	appRouter.ServeHTTP(w, r)
}
