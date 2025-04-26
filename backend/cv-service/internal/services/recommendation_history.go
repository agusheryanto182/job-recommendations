package services

import (
	"cv-service/internal/models"
	"cv-service/internal/repositories"
)

// TODO : this is not final structure
type RecommendationHistoryService interface {
	FindByUserID(ID string) ([]*models.RecommendationHistory, error)
}

type recommendationHistoryService struct {
	recommendationHistoryRepo repositories.RecommendationHistoryRepo
}

func NewRecommendationHistoryService(recommendationHistoryRepo repositories.RecommendationHistoryRepo) RecommendationHistoryService {
	return &recommendationHistoryService{
		recommendationHistoryRepo: recommendationHistoryRepo,
	}
}

func (s *recommendationHistoryService) FindByUserID(ID string) ([]*models.RecommendationHistory, error) {
	return s.recommendationHistoryRepo.FindByUserID(ID)
}
