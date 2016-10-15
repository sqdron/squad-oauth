package oauth

import "time"

type Session struct {
	Code         string
	State        string
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}