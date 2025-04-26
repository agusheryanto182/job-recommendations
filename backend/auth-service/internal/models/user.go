package models

import (
	"auth-service/pkg/uuid"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string         `gorm:"column:id;primary_key" json:"id"`
	GoogleID  string         `gorm:"column:google_id" json:"google_id"`
	Name      string         `gorm:"column:name" json:"name"`
	Email     string         `gorm:"column:email" json:"email"`
	AvatarURL string         `gorm:"column:avatar_url" json:"avatar_url"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.GenerateUUID()
	return nil
}
