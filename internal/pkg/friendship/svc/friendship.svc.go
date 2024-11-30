package svc

import (
	"app/internal/pkg/friendship/ent"
	"context"
	"errors"
	"github.com/google/uuid"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
)

type FriendshipService struct {
	db gossiper.Database
}

func NewFriendshipService(db gossiper.Database) *FriendshipService {
	return &FriendshipService{db: db}
}

func (s *FriendshipService) CreateFriendRequest(ctx context.Context, userID, friendID string) (*ent.Friendship, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}
	friendUUID, err := uuid.Parse(friendID)
	if err != nil {
		return nil, errors.New("invalid friend ID")
	}

	exists := s.db.GetDB().Where("(user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)", userUUID, friendUUID, friendUUID, userUUID).First(&ent.Friendship{})
	if exists.RowsAffected > 0 {
		return nil, errors.New("friend request already exists")
	}

	friendship := &ent.Friendship{
		UserID:   userUUID,
		FriendID: friendUUID,
	}

	if err := s.db.GetDB().Create(friendship).Error; err != nil {
		return nil, err
	}

	return friendship, nil
}

func (s *FriendshipService) AcceptFriendRequest(ctx context.Context, friendshipID string) error {
	friendshipUUID, err := uuid.Parse(friendshipID)
	if err != nil {
		return errors.New("invalid friendship ID")
	}

	var friendship ent.Friendship
	if err := s.db.GetDB().First(&friendship, "id = ?", friendshipUUID).Error; err != nil {
		return errors.New("friend request not found")
	}

	tx := s.db.GetDB().Begin()

	// Add friends relationship
	if err := tx.Exec("INSERT INTO friendships (user_id, friend_id) VALUES (?, ?), (?, ?)",
		friendship.UserID, friendship.FriendID, friendship.FriendID, friendship.UserID).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Remove the request
	if err := tx.Delete(&ent.Friendship{}, friendship.ID).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (s *FriendshipService) GetFriendRequests(ctx context.Context, userID string, inout string) ([]ent.Friendship, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	var friendships []ent.Friendship
	query := s.db.GetDB()
	if inout == "IN" {
		query = query.Where("friend_id = ?", userUUID)
	} else if inout == "OUT" {
		query = query.Where("user_id = ?", userUUID)
	} else {
		return nil, errors.New("invalid inout value")
	}

	if err := query.Find(&friendships).Error; err != nil {
		return nil, err
	}

	return friendships, nil
}

func (s *FriendshipService) RemoveFriendRequest(ctx context.Context, friendshipID string) error {
	friendshipUUID, err := uuid.Parse(friendshipID)
	if err != nil {
		return errors.New("invalid friendship ID")
	}

	if err := s.db.GetDB().Delete(&ent.Friendship{}, "id = ?", friendshipUUID).Error; err != nil {
		return err
	}

	return nil
}
