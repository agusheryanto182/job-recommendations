// pkg/config/config.go
package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	// Server Configuration
	ServerPort string `mapstructure:"PORT"`

	// Logger Configuration
	LogLevel  string `mapstructure:"LOG_LEVEL"`
	LogFile   string `mapstructure:"LOG_FILE"`
	LogFormat string `mapstructure:"LOG_FORMAT"`

	// Database Configuration
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`

	// Google OAuth Configuration
	GoogleClientID     string `mapstructure:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string `mapstructure:"GOOGLE_CLIENT_SECRET"`
	GoogleRedirectURL  string `mapstructure:"GOOGLE_REDIRECT_URL"`

	// JWT Configuration
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

func LoadConfig() (*Config, error) {
	// Set default values
	viper.SetDefault("PORT", "8001")
	viper.SetDefault("LOG_LEVEL", "INFO")
	viper.SetDefault("LOG_FILE", "logs/app.log")
	viper.SetDefault("LOG_FORMAT", "json")

	// Set config file
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	// Read from environment
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	config := &Config{}
	err := viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
