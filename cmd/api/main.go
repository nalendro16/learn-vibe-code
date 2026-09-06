package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nalendro16/learn-vibe-code/internal/config"
	"github.com/nalendro16/learn-vibe-code/internal/database"
	"github.com/nalendro16/learn-vibe-code/internal/server"
	"gorm.io/gorm"
)

func main() {
	// Initialize structured logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	slog.Info("Starting PPOB Service...")

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("Failed to load configuration", "error", err)
		os.Exit(1)
	}

	// Initialize database connection
	var db *gorm.DB
	db, err = database.NewPostgresDB(cfg.DB)
	if err != nil {
		slog.Warn("Could not connect to PostgreSQL database (continuing in development mode)", "error", err)
	}

	// Initialize router
	router := server.NewRouter(cfg, db)

	// Configure HTTP server
	srv := &http.Server{
		Addr:         ":" + cfg.App.Port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in background
	go func() {
		slog.Info("HTTP server is running", "port", cfg.App.Port, "env", cfg.App.Env)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("HTTP server failed to start", "error", err)
			os.Exit(1)
		}
	}()

	// Graceful shutdown handling
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown", "error", err)
		os.Exit(1)
	}

	// Close database connection if initialized
	if db != nil {
		if sqlDB, err := db.DB(); err == nil {
			_ = sqlDB.Close()
		}
	}

	slog.Info("Server exited gracefully")
}
