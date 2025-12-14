package services

import (
	"ginLearning/05_Auth/models"

	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func InitAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		db: db,
	}
}

func (a *AuthService) LoginService(email *string, password *string) models.Login {

}
