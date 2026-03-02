package user

import (
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

func (r *Repository) CreateUser(user *User) (*User, error) {
	err := r.db.Create(user).Error
	return user, err
}
