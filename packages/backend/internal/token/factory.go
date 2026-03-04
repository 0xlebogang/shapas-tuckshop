package token

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type Factory struct {
	secretKey []byte
}

func NewFactory(secretKey string) *Factory {
	return &Factory{
		[]byte(secretKey),
	}
}

func (f *Factory) CreateToken(userID, sessionID string, duration time.Duration) (string, error) {
	claims := UserClaims{
		UserID:    userID,
		SessionID: sessionID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(f.secretKey)
}

func (g *Factory) VerifyToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return g.secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	return claims, nil
}
