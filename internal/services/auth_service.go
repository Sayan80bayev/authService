package service

import (
	"authService/internal/models"
	"authService/internal/repository"
	"errors"

	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// AuthService — сервис для работы с аутентификацией пользователей.
// Содержит репозиторий пользователей и ключ для подписи JWT токенов.
type AuthService struct {
	userRepo *repository.UserRepository
	jwtKey   []byte
}

// NewAuthService создает новый экземпляр AuthService
func NewAuthService(userRepo *repository.UserRepository, jwtKey string) *AuthService {
	return &AuthService{userRepo, []byte(jwtKey)}
}

// Register регистрирует нового пользователя:
// 1) Хэширует пароль
// 2) Сохраняет пользователя в базе
// 3) Генерирует JWT токен для авторизованного доступа
func (uc *AuthService) Register(email, password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	// Создаем объект пользователя
	user := &models.User{
		Email:    email,
		Password: string(hashedPassword),
	}
	// Сохраняем пользователя в базе
	err = uc.userRepo.CreateUser(user)
	if err != nil {
		return "", err
	}
	// Генерируем JWT токен для нового пользователя
	token := generateJwtToken(user)

	return token.SignedString(uc.jwtKey)
}

// Login выполняет авторизацию пользователя:
// 1) Проверяет наличие пользователя по email
// 2) Сравнивает хэш пароля с введенным паролем
// 3) Возвращает JWT токен, если данные верные
func (uc *AuthService) Login(email, password string) (string, error) {
	// Получаем пользователя по email
	user, err := uc.userRepo.GetUserByUsername(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	// Проверяем совпадает ли пароль с хэшем
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token := generateJwtToken(user)

	return token.SignedString(uc.jwtKey)
}

// generateJwtToken создает новый JWT токен с данными пользователя:
// user_id, email, роль, статус активности, срок жизни (24 часа)
func generateJwtToken(user *models.User) *jwt.Token {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":     user.ID,
		"email":       user.Email,
		"user_role":   user.Role,
		"user_active": user.Active,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	})
	return token
}
