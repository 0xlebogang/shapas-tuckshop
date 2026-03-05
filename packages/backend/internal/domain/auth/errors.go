package auth

import "errors"

var (
	ErrInvalidToken   = errors.New("Invalid refresh token")
	ErrSessionExpired = errors.New("Session has expired or been revoked")
)
