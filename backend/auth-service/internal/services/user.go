package services

import (
	"auth-service/internal/errs"
	"auth-service/internal/models"
	"auth-service/internal/repositories"
	"auth-service/internal/request"
	"auth-service/internal/response"
	"encoding/json"
	"io"
	"net/http"
)

type AuthService interface {
	AuthenticateGoogle(authenticateUserRequest request.AuthenticateUserRequest) (*models.User, error)
	GetUserProfile(ID string) (*models.User, error)
	GetGoogleUserInfo(accessToken string) (*response.GoogleUserInfo, error)
	ValidateGoogleState(state, savedState string) error
}

type authService struct {
	userRepo repositories.UserRepo
}

func NewAuthService(userRepo repositories.UserRepo) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (s *authService) ValidateGoogleState(state, savedState string) error {
	if state != savedState {
		return errs.NewUnprocessableEntityError("Invalid state")
	}
	return nil
}

func (s *authService) GetGoogleUserInfo(accessToken string) (*response.GoogleUserInfo, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + accessToken)
	if err != nil {
		return nil, errs.NewInternalServerError("Failed to get user info")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errs.NewInternalServerError("Failed to read response body")
	}

	var userInfo response.GoogleUserInfo
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, errs.NewInternalServerError("Failed to unmarshal response body")
	}

	// Validate data
	if !userInfo.VerifiedEmail {
		return nil, errs.NewUnprocessableEntityError("Email is not verified")
	}
	if userInfo.Email == "" || userInfo.ID == "" {
		return nil, errs.NewUnprocessableEntityError("Invalid user data from Google")
	}

	return &userInfo, nil
}

func (s *authService) GetUserProfile(ID string) (*models.User, error) {
	return s.userRepo.FindByID(ID)
}

func (s *authService) AuthenticateGoogle(authenticateUserRequest request.AuthenticateUserRequest) (*models.User, error) {
	user, _ := s.userRepo.FindByGoogleID(authenticateUserRequest.ID)
	if user != nil {
		return user, nil
	}

	user = &models.User{
		GoogleID:  authenticateUserRequest.ID,
		Name:      authenticateUserRequest.Name,
		Email:     authenticateUserRequest.Email,
		AvatarURL: authenticateUserRequest.Picture,
	}

	err := s.userRepo.Create(user)
	if err != nil {
		return nil, errs.NewInternalServerError("Failed to create user")
	}

	return user, err
}
