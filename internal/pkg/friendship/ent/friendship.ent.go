package ent

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Friendship struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	FriendID  uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// BeforeCreate Hook for generating custom UUID for friendship ID
func (f *Friendship) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate custom UUID for friendship
	f.ID = uuid.New()

	return nil
}
