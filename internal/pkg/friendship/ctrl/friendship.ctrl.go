package ctrl

import (
	pbu "app/internal/core/grpc/generated/generic/utils"
	pbFriendship "app/internal/core/grpc/generated/lotof.hub.msvc.users/friendship"
	pbUser "app/internal/core/grpc/generated/lotof.hub.msvc.users/user"
	"app/internal/pkg/friendship/ent"
	"app/internal/pkg/friendship/svc"
	"context"
	"log"
)

type FriendshipController struct {
	friendshipService *svc.FriendshipService
	pbFriendship.UnimplementedFriendshipServiceServer
}

func NewFriendshipController(service *svc.FriendshipService) *FriendshipController {
	return &FriendshipController{
		friendshipService: service,
	}
}

func (f FriendshipController) CreateFriendship(_ context.Context, input *pbFriendship.CreateFriendshipInput) (*pbFriendship.Friendship, error) {
	friendship, err := f.friendshipService.CreateFriendRequest(input.UserId, input.FriendId)

	if err != nil {
		return nil, err
	}
	return &pbFriendship.Friendship{
		Id: friendship.ID.String(),
		User: &pbUser.User{
			Id:       friendship.User.ID.String(),
			Username: friendship.User.Username,
			Email:    friendship.User.Email,
		},
		Friend: &pbUser.User{
			Id:       friendship.Friend.ID.String(),
			Username: friendship.Friend.Username,
			Email:    friendship.Friend.Email,
		},
		Status: uint32(friendship.Status),
	}, err
}

func (f FriendshipController) AcceptFriendshipRequest(_ context.Context, input *pbFriendship.AcceptFriendshipInput) (*pbFriendship.Friendship, error) {
	accept, err := f.friendshipService.AcceptFriendRequest(input.FriendshipId)
	if err != nil {
		return nil, err
	}
	return &pbFriendship.Friendship{
		Id: accept.ID.String(),
		User: &pbUser.User{
			Id:       accept.User.ID.String(),
			Username: accept.User.Username,
			Email:    accept.User.Email,
		},
		Friend: &pbUser.User{
			Id:       accept.Friend.ID.String(),
			Username: accept.Friend.Username,
			Email:    accept.Friend.Email,
		},
		Status: uint32(accept.Status),
	}, nil
}

func (f FriendshipController) RemoveFriendshipRequest(_ context.Context, input *pbFriendship.RemoveFriendshipInput) (*pbu.Empty, error) {
	err := f.friendshipService.RemoveFriend(input.FriendshipId)
	if err != nil {
		log.Printf("Error removing friendship: %v", err)
		return nil, err
	}

	return &pbu.Empty{}, nil
}

func (f FriendshipController) FriendshipRequestList(_ context.Context, filter *pbFriendship.FriendshipFilter) (*pbFriendship.PaginatedFriendshipList, error) {
	friendships, err := f.friendshipService.GetFriendships(filter.UserId, ent.FriendshipStatus(filter.Status))
	if err != nil {
		return nil, err
	}

	var friendshipList []*pbFriendship.Friendship
	for _, friendship := range friendships {
		friendshipList = append(friendshipList, &pbFriendship.Friendship{
			Id: friendship.ID.String(),
			User: &pbUser.User{
				Id:       friendship.User.ID.String(),
				Username: friendship.User.Username,
				Email:    friendship.User.Email,
			},
			Friend: &pbUser.User{
				Id:       friendship.Friend.ID.String(),
				Username: friendship.Friend.Username,
				Email:    friendship.Friend.Email,
			},
			Status: uint32(friendship.Status),
		})
	}

	return &pbFriendship.PaginatedFriendshipList{
		Friendships: friendshipList,
		PaginationInfo: &pbu.PaginationInfo{
			Count: int32(len(friendshipList)),
		},
	}, nil
}
