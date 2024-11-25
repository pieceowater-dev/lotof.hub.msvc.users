package svc

import (
	"app/internal/core/db"
	"app/internal/pkg/user/ent"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService struct {
	db db.Database
}

func NewAuthService(db db.Database) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (*ent.User, error) {
	var user ent.User
	if err := s.db.GetDB().Where("email = ? AND deleted = ?", email, false).First(&user).Error; err != nil {
		return nil, errors.New("incorrect user or password")
	}

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("incorrect user or password")
	}

	return &user, nil
}

func (s *AuthService) Register(ctx context.Context, user *ent.User) (*ent.User, error) {
	// Hash the password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()

	if err := s.db.GetDB().Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
