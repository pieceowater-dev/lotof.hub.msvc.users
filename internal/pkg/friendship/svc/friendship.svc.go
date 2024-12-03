package svc

import (
	"app/internal/pkg/friendship/ent"
	"errors"
	"github.com/google/uuid"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"log"
)

type FriendshipService struct {
	db gossiper.Database
}

func NewFriendshipService(db gossiper.Database) *FriendshipService {
	return &FriendshipService{db: db}
}

func (s *FriendshipService) CreateFriendRequest(userID, friendID string) (*ent.Friendship, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}
	friendUUID, err := uuid.Parse(friendID)
	if err != nil {
		return nil, errors.New("invalid friend ID")
	}

	var existingFriendship ent.Friendship
	exists := s.db.GetDB().
		Where("(user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)",
			userUUID, friendUUID, friendUUID, userUUID).
		First(&existingFriendship)
	if exists.RowsAffected > 0 {
		return nil, errors.New("friend request already exists")
	}

	friendship := &ent.Friendship{
		UserID:   userUUID,
		FriendID: friendUUID,
		Status:   ent.Pending,
	}

	if err := s.db.GetDB().Create(friendship).Error; err != nil {
		return nil, err
	}

	if err := s.db.GetDB().
		Preload("User").
		Preload("Friend").
		First(friendship, "id = ?", friendship.ID).
		Error; err != nil {
		return nil, err
	}

	return friendship, nil
}

func (s *FriendshipService) GetFriendship(friendshipUUID uuid.UUID) (*ent.Friendship, error) {
	var friendship ent.Friendship
	if err := s.db.GetDB().
		Preload("User").
		Preload("Friend").
		First(&friendship, "id = ?", friendshipUUID).Error; err != nil {
		return nil, errors.New("friendship not found with includes")
	}

	return &friendship, nil
}

func (s *FriendshipService) AcceptFriendRequest(friendshipID string) (*ent.Friendship, error) {
	friendshipUUID, err := uuid.Parse(friendshipID)
	if err != nil {
		return nil, errors.New("invalid friendship ID")
	}

	var friendship ent.Friendship
	if err := s.db.GetDB().
		First(&friendship, "id = ?", friendshipUUID).Error; err != nil {
		return nil, errors.New("friend request not found")
	}

	friendship.Status = ent.Accepted
	if err := s.db.GetDB().Save(&friendship).Error; err != nil {
		log.Println("Error while saving:", err)
		return nil, err
	}

	return s.GetFriendship(friendshipUUID)
}

func (s *FriendshipService) RemoveFriend(friendshipID string) error {
	friendshipUUID, err := uuid.Parse(friendshipID)
	if err != nil {
		return errors.New("invalid friendship ID")
	}

	if err := s.db.GetDB().Delete(&ent.Friendship{}, "id = ?", friendshipUUID).Error; err != nil {
		return err
	}

	return nil
}

func (s *FriendshipService) GetFriendships(userID string, status ent.FriendshipStatus) ([]ent.Friendship, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	var friendships []ent.Friendship

	if err := s.db.GetDB().
		Where("user_id = ? OR friend_id = ?", userUUID, userUUID).
		Where("status = ?", status).
		Preload("User").
		Preload("Friend").
		Find(&friendships).Error; err != nil {
		return nil, err
	}

	return friendships, nil
}
