package auth

import (
	"time"

	"github.com/0xlebogang/shapas/internal/domain/user"
	"github.com/0xlebogang/shapas/internal/token"
	"github.com/0xlebogang/shapas/internal/util"
)

type Service struct {
	factory *token.Factory
	repo    *user.Repository
}

func NewService(factory *token.Factory) *Service {
	return &Service{factory: factory}
}

func (s *Service) Authenticate(json *RequestBody) (*map[string]string, error) {
	user, err := s.repo.GetUserByEmail(json.Email)
	if err != nil {
		return nil, err
	}

	if err := util.CheckPassword(json.Password, *user.Password); err != nil {
		return nil, err
	}

	var tokens map[string]string

	accessToken, err := s.factory.CreateToken(user.ID, user.Email, *user.Role, 60*time.Minute)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.factory.CreateToken(user.ID, user.Email, *user.Role, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	tokens["access_token"] = accessToken
	tokens["refresh_token"] = refreshToken

	return &tokens, nil
}

func (s *Service) VerifyToken(tokenString string) (*token.UserClaims, error) {
	return s.factory.VerifyToken(tokenString)
}
