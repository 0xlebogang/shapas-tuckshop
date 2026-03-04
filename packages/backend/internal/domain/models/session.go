package models

import (
	"time"
)

type Session struct {
	ID           string    `json:"id" gorm:"primaryKey;type:varchar(50)"`
	UpdatedAt    time.Time `json:"-" gorm:"-"`
	UserID       string    `json:"user_id"`
	RefreshToken string    `json:"refresh_token" gorm:"not null"`
	IsRevoked    bool      `json:"is_revoked" gorm:"default:false"`
	ExpiresAt    time.Time `json:"expires_at" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at"`

	User *User `json:"user,omitempty" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
