package models

import (
	"cv-service/pkg/uuid"
	"time"

	"gorm.io/gorm"
)

type RecommendationHistory struct {
	ID              string              `gorm:"column:id;primary_key" json:"id"`
	UserID          string              `gorm:"column:user_id" json:"user_id"`
	RecommendedJobs []JobRecommendation `gorm:"column:recommended_jobs" json:"recommended_jobs"`
	CreatedAt       time.Time           `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time           `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt       gorm.DeletedAt      `gorm:"column:deleted_at" json:"deleted_at"`
}

// TODO : this is not the final structure yet
type JobRecommendation struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	CompanyName   string    `json:"company_name"`
	Description   string    `json:"description"`
	Location      string    `json:"location"`
	SalaryMin     float64   `json:"salary_min"`
	SalaryMax     float64   `json:"salary_max"`
	JobType       string    `json:"job_type"`
	Score         float64   `json:"score"`
	MatchedSkills []Skill   `json:"matched_skills"`
	RecommendedAt time.Time `json:"recommended_at"`
}

// TODO : this is not the final structure yet
type Skill struct {
	Name   string  `json:"skill"`
	Weight float64 `json:"weight"`
}

func (RecommendationHistory) TableName() string {
	return "recommendation_history"
}

func (u *RecommendationHistory) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.GenerateUUID()
	return nil
}
