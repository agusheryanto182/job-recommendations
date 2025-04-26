package repositories

import (
	"auth-service/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type InvalidTokenRepository struct {
	db *gorm.DB
}

func NewInvalidTokenRepository(db *gorm.DB) InvalidTokenRepository {
	return InvalidTokenRepository{db}
}

func (r InvalidTokenRepository) FindByToken(token string) (*models.InvalidToken, error) {
	var invalidToken models.InvalidToken
	err := r.db.Debug().Where("token = ?", token).First(&invalidToken).Error

	return &invalidToken, err
}

func (r InvalidTokenRepository) Create(newInvalidToken models.InvalidToken) (models.InvalidToken, error) {
	err := r.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&newInvalidToken).Error

	return newInvalidToken, err
}
