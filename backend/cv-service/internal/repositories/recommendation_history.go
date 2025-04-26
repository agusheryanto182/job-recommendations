package repositories

import (
	"cv-service/internal/models"

	"gorm.io/gorm"
)

type RecommendationHistoryRepo interface {
	FindByUserID(ID string) ([]*models.RecommendationHistory, error)
}

type recommendationHistoryRepo struct {
	db *gorm.DB
}

func NewRecommendationHistoryRepo(db *gorm.DB) RecommendationHistoryRepo {
	return &recommendationHistoryRepo{
		db: db,
	}
}

func (r *recommendationHistoryRepo) FindByUserID(ID string) ([]*models.RecommendationHistory, error) {
	var recommendationHistory []*models.RecommendationHistory
	err := r.db.Debug().Where("user_id = ?", ID).Find(&recommendationHistory).Error

	return recommendationHistory, err
}
