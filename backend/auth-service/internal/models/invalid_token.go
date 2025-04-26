package models

import (
	"time"
)

const TableNameInvalidToken = "invalid_tokens"

type InvalidToken struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Token     string    `gorm:"column:token;not null" json:"token"`
	ExpiredAt time.Time `gorm:"column:expired_at;not null" json:"expired_at"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (*InvalidToken) TableName() string {
	return TableNameInvalidToken
}
