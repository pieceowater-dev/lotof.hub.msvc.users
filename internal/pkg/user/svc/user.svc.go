package svc

import (
	"app/internal/core/db"
	"app/internal/pkg/user/ent"
	"errors"
	"gorm.io/gorm"
)

type UserService struct {
	db db.Database
}

func NewUserService(db db.Database) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(user *ent.User) (*ent.User, error) {
	if err := s.db.GetDB().Create(user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errors.New("email already exists")
		}
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUserByID(id string) (*ent.User, error) {
	var user ent.User
	if err := s.db.GetDB().Preload("Friends").First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) UpdateUser(user *ent.User) (*ent.User, error) {
	if err := s.db.GetDB().Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) DeleteUser(id string) error {
	return s.db.GetDB().Delete(&ent.User{}, "id = ?", id).Error
}

func (s *UserService) AddFriend(userID, friendID string) error {
	var user, friend ent.User
	if err := s.db.GetDB().First(&user, "id = ?", userID).Error; err != nil {
		return err
	}
	if err := s.db.GetDB().First(&friend, "id = ?", friendID).Error; err != nil {
		return err
	}
	return s.db.GetDB().Model(&user).Association("Friends").Append(&friend)
}
