package token

import jwt "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	UserID    string `json:"user_id"`
	SessionID string `json:"session_id,omitempty"`
	jwt.RegisteredClaims
}
