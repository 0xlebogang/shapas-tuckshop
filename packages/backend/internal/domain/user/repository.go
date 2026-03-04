package user

import (
	"github.com/0xlebogang/shapas/internal/domain/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateUser(value *models.User) (*models.User, error) {
	err := r.db.Create(&value).Error
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (r *Repository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
