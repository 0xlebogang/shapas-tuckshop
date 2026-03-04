package user

import (
	"time"

	"github.com/0xlebogang/shapas/internal/domain/common"
)

type User struct {
	common.BaseModel
	Email    string  `json:"email" gorm:"type:varchar(255);uniqueIndex;not null" binding:"required,email,max=255"`
	Name     *string `json:"name" gorm:"type:varchar(255)" binding:"omitempty,min=2,max=255"`
	Avatar   *string `json:"avatar" gorm:"type:text" binding:"omitempty,url"`
	Password *string `json:"password" gorm:"type:text" binding:"omitempty,min=6"`
	Role     *string `json:"role" gorm:"type:varchar(50)" binding:"omitempty,min=2,max=50"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      *string   `json:"name"`
	Avatar    *string   `json:"avatar"`
	Role      *string   `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		Email:     u.Email,
		Name:      u.Name,
		Avatar:    u.Avatar,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
