package main

import (
	"log/slog"

	"github.com/ahmed/shophub-backend/infra"
	"github.com/gin-gonic/gin"
)

func main() {

	if err := infra.InitEnv(); err != nil {
		slog.Error("Failed to initialize environment", "error", err)
		return
	}
	slog.Info("Environment initialized successfully")
	slog.Info("Connecting Database")

	if err := infra.DBConnection(); err != nil {
		slog.Error("Failed to connect database", "error", err)
		return
	}
	slog.Info("Database connected successfully")

	r := gin.Default()

	if err := r.Run(":" + infra.Configuration.Port); err != nil {
		slog.Error("Failed to start server", "error", err)
		return
	}
}
