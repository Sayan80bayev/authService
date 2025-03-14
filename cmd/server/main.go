package main

import (
	"authService/internal/config"
	"authService/internal/routes"
	"authService/pkg/logging"
	"gorm.io/driver/postgres"

	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	r := gin.Default()
	r.Use(gin.LoggerWithFormatter(logging.CustomLogFormatter), gin.Recovery())
	routes.SetupAuthRoutes(r, db, cfg)

	err = r.Run(":" + cfg.Port)
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}

}
