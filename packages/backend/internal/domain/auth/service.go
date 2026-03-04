package auth

import (
	"time"

	"github.com/0xlebogang/shapas/internal/domain/user"
	"github.com/0xlebogang/shapas/internal/token"
	"github.com/0xlebogang/shapas/internal/util"
)

type Service struct {
	tokenFactory *token.Factory
	repo         *user.Repository
}

func NewService(factory *token.Factory, repo *user.Repository) *Service {
	return &Service{tokenFactory: factory, repo: repo}
}

func (s *Service) Authenticate(json *RequestBody) (*AuthTokens, error) {
	user, err := s.repo.GetUserByEmail(json.Email)
	if err != nil {
		return nil, err
	}

	if err := util.CheckPassword(json.Password, *user.Password); err != nil {
		return nil, err
	}

	accessToken, err := s.tokenFactory.CreateToken(user.ID, user.Email, user.Role, 60*time.Minute)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.tokenFactory.CreateToken(user.ID, user.Email, user.Role, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	tokens := AuthTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &tokens, nil
}

func (s *Service) VerifyToken(tokenString string) (*token.UserClaims, error) {
	return s.tokenFactory.VerifyToken(tokenString)
}
