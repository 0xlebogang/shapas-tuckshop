package auth

import (
	"time"

	"github.com/0xlebogang/shapas/internal/domain/models"
	"github.com/0xlebogang/shapas/internal/domain/user"
	"github.com/0xlebogang/shapas/internal/token"
	"github.com/0xlebogang/shapas/internal/util"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Service struct {
	tokenFactory *token.Factory
	userRepo     *user.Repository
	authRepo     *Repository
}

func NewService(factory *token.Factory, userRepo *user.Repository, authRepo *Repository) *Service {
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

	accessToken, err := s.tokenFactory.CreateToken(user.ID, "", 60*time.Minute)
	if err != nil {
		return nil, err
	}

	sessionID, _ := gonanoid.New()

	refreshToken, err := s.tokenFactory.CreateToken(user.ID, sessionID, 7*24*time.Hour)
	if err != nil {
		return nil, err
	}

	s.authRepo.CreateSession(&models.Session{
		ID:           sessionID,
		UserID:       user.ID,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(24 * time.Hour),
	})

	tokens := AuthTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &tokens, nil
}

func (s *Service) VerifyToken(tokenString string) (*token.UserClaims, error) {
	return s.tokenFactory.VerifyToken(tokenString)
}
