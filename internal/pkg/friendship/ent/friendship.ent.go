package ent

import (
	entUser "app/internal/pkg/user/ent"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FriendshipStatus int

const (
	Pending  FriendshipStatus = 100
	Accepted FriendshipStatus = 200
	Rejected FriendshipStatus = 300
)

type Friendship struct {
	gorm.Model
	ID       uuid.UUID        `gorm:"type:uuid;primaryKey"`
	UserID   uuid.UUID        `gorm:"type:uuid;not null"`
	FriendID uuid.UUID        `gorm:"type:uuid;not null"`
	Status   FriendshipStatus `gorm:"type:smallint;default:100"` // Default to Pending

	// Связи с таблицей пользователей
	User   *entUser.User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Friend *entUser.User `gorm:"foreignKey:FriendID;constraint:OnDelete:CASCADE"`
}

// BeforeCreate Hook for generating custom UUID for friendship ID
func (f *Friendship) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate custom UUID for friendship
	f.ID = uuid.New()

	return nil
}
