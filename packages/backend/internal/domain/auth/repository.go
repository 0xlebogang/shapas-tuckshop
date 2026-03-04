package auth

import (
	"github.com/0xlebogang/shapas/internal/domain/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateSession(session *models.Session) error {
	return r.db.Create(&session).Error
}
