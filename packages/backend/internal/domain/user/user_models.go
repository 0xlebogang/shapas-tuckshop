package user

import "github.com/0xlebogang/shapas/backend/internal/domain/common"

type User struct {
	common.BaseModel
	Email    string  `json:"email" gorm:"type:varchar(255);uniqueIndex;not null" binding:"required;max=255"`
	Name     *string `json:"name" gorm:"type:varchar(255)" binding:"omitempty;min=2;max=255"`
	Avatar   *string `json:"avatar" gorm:"type:text" binding:"omitempty;url"`
	Password *string `json:"omitempty;password" gorm:"type:text;->:false;<-:create;<-:update" binding:"omitempty;min=6"`
}
