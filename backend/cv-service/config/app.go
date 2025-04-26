package config

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Initialize() (*gorm.DB, error) {
	cfg, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	logrus.Info("Initializing database connection...")

	return ConnectDB(cfg)
}
