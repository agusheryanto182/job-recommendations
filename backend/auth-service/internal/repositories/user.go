package repositories

import (
	"auth-service/internal/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	FindByGoogleID(googleID string) (*models.User, error)
	Create(user *models.User) error
	FindByID(ID string) (*models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) FindByGoogleID(googleID string) (*models.User, error) {
	var user models.User
	err := r.db.Where("google_id = ?", googleID).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) FindByID(ID string) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", ID).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
