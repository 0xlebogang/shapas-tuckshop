package auth

import (
	"time"

	"github.com/0xlebogang/shapas/internal/constants"
	"github.com/0xlebogang/shapas/internal/domain/models"
	"github.com/0xlebogang/shapas/internal/domain/user"
	"github.com/0xlebogang/shapas/internal/token"
	"github.com/0xlebogang/shapas/internal/util"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Service struct {
	tokenFactory *token.TokenFactory
	userRepo     *user.Repository
	authRepo     *Repository
}

func NewService(factory *token.TokenFactory, userRepo *user.Repository, authRepo *Repository) *Service {
	return &Service{tokenFactory: factory, userRepo: userRepo, authRepo: authRepo}
}

func (s *Service) Authenticate(json *AuthRequest) (*AuthTokens, error) {
	user, err := s.userRepo.GetUserByEmail(json.Email)
	if err != nil {
		return nil, err
	}

	if err := util.CheckPassword(json.Password, *user.Password); err != nil {
		return nil, err
	}

	tokens, err := s.generateTokens(user.ID)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (s *Service) Refresh(json *RefreshRequest) (*AuthTokens, error) {
	claims, err := s.tokenFactory.VerifyToken(json.RefreshToken)
	if err != nil {
		return nil, ErrInvalidToken
	}

	session, err := s.authRepo.GetSession(claims.SessionID, claims.UserID)
	if err != nil {
		return nil, err
	}

	if session.IsRevoked || time.Now().After(session.ExpiresAt) {
		return nil, ErrSessionExpired
	}

	tokens, err := s.generateTokens(claims.UserID)
	if err != nil {
		return nil, err
	}

	if err := s.authRepo.DeleteSession(claims.SessionID, claims.UserID); err != nil {
		return nil, err
	}

	return tokens, nil
}

func (s *Service) generateTokens(userID string) (*AuthTokens, error) {
	accessToken, err := s.tokenFactory.CreateToken(userID, "", constants.AuthAccessTokenLifeTime)
	if err != nil {
		return nil, err
	}

	sessionID, _ := gonanoid.New()
	refreshToken, err := s.tokenFactory.CreateToken(userID, sessionID, constants.AuthRefreshTokenLifeTime)
	if err != nil {
		return nil, err
	}

	tokens := AuthTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	if err := s.authRepo.CreateSession(&models.Session{
		ID:           sessionID,
		UserID:       userID,
		RefreshToken: tokens.RefreshToken,
		ExpiresAt:    time.Now().Add(constants.AuthRefreshTokenLifeTime),
	}); err != nil {
		return nil, err
	}

	return &tokens, nil
}
