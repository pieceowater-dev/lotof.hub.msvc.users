package ent

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserState int

const (
	Suspended UserState = 100 // waiting for activation
	Active    UserState = 200 // active user, ok
	Blocked   UserState = 400 // blocked user
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Username string    `gorm:"type:varchar(255);unique;not null"`
	Email    string    `gorm:"type:varchar(255);unique;not null"`
	Password string    `gorm:"type:varchar(255);not null"`
	State    UserState `gorm:"type:smallint;default:100"` // Default to Suspended
	Friends  []*User   `gorm:"many2many:friendships;joinForeignKey:UserID;joinReferences:FriendID"`
}

// BeforeCreate Hook for generating UUID
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return nil
}
