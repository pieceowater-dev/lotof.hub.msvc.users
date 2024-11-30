package ent

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserState int

const (
	Suspended UserState = 100 // waiting for activation
	Active    UserState = 200 // active user, ok
	Blocked   UserState = 500 // blocked user
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Username  string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"type:varchar(255);unique;not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	State     UserState `gorm:"default:100"` // Default to Suspended
	Deleted   bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Friends   []*User   `gorm:"many2many:friendships;joinForeignKey:UserID;joinReferences:FriendID"`
}

// BeforeCreate Hook for generating custom UUID and password hashing
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate custom UUID for user
	u.ID = uuid.New()

	// Hash password if it's not empty
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return nil
}

// BeforeSave Hook for updating timestamp and hashing password (if necessary)
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return nil
}
