package models

import (
	"cv-service/pkg/uuid"
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type RecommendationHistory struct {
	ID              string         `gorm:"column:id;primary_key" json:"id"`
	UserID          string         `gorm:"column:user_id" json:"user_id"`
	RecommendedJobs JobsJSON       `gorm:"column:recommended_jobs;type:json" json:"recommended_jobs"`
	CreatedAt       time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

type JobsJSON []JobRecommendation

func (j JobsJSON) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JobsJSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, j)
}

type JobRecommendation struct {
	ID            string     `json:"id"`
	Title         string     `json:"title"`
	CompanyName   string     `json:"company_name"`
	Description   string     `json:"description"`
	Location      string     `json:"location"`
	Salary        string     `json:"salary"`
	JobType       string     `json:"job_type"`
	Score         float64    `json:"score"`
	MatchedSkills SkillsJSON `json:"matched_skills"`
	NumbApplicant string     `json:"numb_applicant"`
	URL           string     `json:"url"`
	PostedDate    string     `json:"posted_date"`
}

type SkillsJSON []Skill

func (s SkillsJSON) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *SkillsJSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, &s)
}

type Skill struct {
	Skill string `json:"skill"`
}

func (RecommendationHistory) TableName() string {
	return "recommendation_history"
}

func (u *RecommendationHistory) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.GenerateUUID()
	return nil
}
