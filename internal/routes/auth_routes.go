package routes

import (
	"authService/internal/config"
	"authService/internal/delivery"
	"authService/internal/repository"
	"authService/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAuthRoutes(r *gin.Engine, db *gorm.DB, cfg *config.Config) {
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo, cfg.JWTSecret)
	authHandler := delivery.NewAuthHandler(authService)

	// Роуты для аутентификации
	r.POST("/api/v1/register", authHandler.Register)
	r.POST("/api/v1/auth", authHandler.Login)
}
