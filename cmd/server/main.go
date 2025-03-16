package main

import (
	"authService/internal/config"
	"authService/internal/routes"
	"authService/pkg/logging"
	"gorm.io/driver/postgres"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var logger = logging.GetLogger()

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("Error loading config:", err)
		return
	}

	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		logger.Error("Ошибка подключения к базе данных:", err)
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
