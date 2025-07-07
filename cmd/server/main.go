package main

import (
	"authService/internal/config"
	"authService/internal/models" // 👈 Import your models here
	"authService/internal/routes"
	"authService/pkg/logging"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var logger = logging.GetLogger()

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("Error loading config:", err)
		return
	}

	logger.Info(cfg.DatabaseURL)

	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		logger.Error("Could not connect to db:", err)
		return
	}

	if err := db.AutoMigrate(
		&models.User{},
	); err != nil {
		logger.Error("Failed to run migrations:", err)
		return
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(logging.Middleware)

	routes.SetupAuthRoutes(r, db, cfg)

	logger.Info("All needed connections are made")
	logger.Infof("🚀 Server is running on port :%s", cfg.Port)

	err = r.Run(":" + cfg.Port)
	if err != nil {
		logger.Error("Failed to start server:", err)
		return
	}
}
