package oauth

import "time"

type Session struct {
	ID           string
	AuthURL      string
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}