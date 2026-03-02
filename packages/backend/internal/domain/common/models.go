package common

import (
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string     `json:"id" gorm:"primaryKey;type:varchar(21);<-:create"`
	CreatedAt time.Time  `json:"created_at" gorm:"<-:create"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"<-:update"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	if b.ID == "" {
		id, err := gonanoid.New()
		if err != nil {
			return err
		}
		b.ID = id
	}
	return nil
}
