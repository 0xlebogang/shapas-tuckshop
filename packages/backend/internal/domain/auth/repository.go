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

func (r *Repository) GetSession(sessionID, userID string) (*models.Session, error) {
	var session models.Session
	if err := r.db.Where("id = ? AND user_id = ?", sessionID, userID).First(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *Repository) DeleteSession(sessionID, userID string) error {
	session, err := r.GetSession(sessionID, userID)
	if err != nil {
		return err
	}
	return r.db.Delete(&session).Error
}
