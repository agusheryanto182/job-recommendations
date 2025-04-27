// config/google_oauth.go
package config

import (
	"errors"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleOAuthConfig struct {
	config *oauth2.Config
}

// InitGoogleAuth initializes Google OAuth configuration
func InitGoogleAuth(cfg *Config) (*GoogleOAuthConfig, error) {
	if cfg.GoogleClientID == "" {
		return nil, errors.New("GOOGLE_CLIENT_ID is not set")
	}
	if cfg.GoogleClientSecret == "" {
		return nil, errors.New("GOOGLE_CLIENT_SECRET is not set")
	}
	if cfg.GoogleRedirectURL == "" {
		return nil, errors.New("GOOGLE_REDIRECT_URL is not set")
	}

	config := &oauth2.Config{
		ClientID:     cfg.GoogleClientID,
		ClientSecret: cfg.GoogleClientSecret,
		RedirectURL:  cfg.GoogleRedirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	return &GoogleOAuthConfig{config: config}, nil
}

// GetConfig returns the OAuth configuration
func (g *GoogleOAuthConfig) GetConfig() *oauth2.Config {
	return g.config
}
