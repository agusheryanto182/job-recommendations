package services

import (
	"auth-service/internal/models"
	"auth-service/internal/repositories"
)

type AuthService interface {
	AuthenticateGoogle(googleID, email, name, avatarURL string) (*models.User, error)
	GetUserProfile(googleID string) (*models.User, error)
}

type authService struct {
	userRepo repositories.UserRepo
}

func NewAuthService(userRepo repositories.UserRepo) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (s *authService) GetUserProfile(googleID string) (*models.User, error) {
	return s.userRepo.FindByGoogleID(googleID)
}

func (s *authService) AuthenticateGoogle(googleID, email, name, avatarURL string) (*models.User, error) {
	user, err := s.userRepo.FindByGoogleID(googleID)
	if err == nil && user != nil {
		return user, nil
	}

	user = &models.User{
		GoogleID:  googleID,
		Name:      name,
		Email:     email,
		AvatarURL: avatarURL,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, err
}
