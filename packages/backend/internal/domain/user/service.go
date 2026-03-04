package user

import (
	"github.com/0xlebogang/shapas/internal/domain/models"
	"github.com/0xlebogang/shapas/internal/util"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateUser(value *models.User) (*models.User, error) {
	hashedPassword, err := util.HashPassword(*value.Password)
	if err != nil {
		return nil, err
	}
	value.Password = &hashedPassword
	return s.repo.CreateUser(value)
}
