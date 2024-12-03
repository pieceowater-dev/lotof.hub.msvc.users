package ctrl

import (
	pb "app/internal/core/grpc/generated"
	"app/internal/pkg/friendship/ent"
	"app/internal/pkg/friendship/svc"
	"context"
	"log"
)

type FriendshipController struct {
	friendshipService *svc.FriendshipService
	pb.UnimplementedFriendshipServiceServer
}

func NewFriendshipController(service *svc.FriendshipService) *FriendshipController {
	return &FriendshipController{
		friendshipService: service,
	}
}

func (f FriendshipController) CreateFriendship(_ context.Context, input *pb.CreateFriendshipInput) (*pb.Friendship, error) {
	friendship, err := f.friendshipService.CreateFriendRequest(input.UserId, input.FriendId)

	if err != nil {
		return nil, err
	}
	return &pb.Friendship{
		Id: friendship.ID.String(),
		User: &pb.User{
			Id:       friendship.User.ID.String(),
			Username: friendship.User.Username,
			Email:    friendship.User.Email,
		},
		Friend: &pb.User{
			Id:       friendship.Friend.ID.String(),
			Username: friendship.Friend.Username,
			Email:    friendship.Friend.Email,
		},
		Status: uint32(friendship.Status),
	}, err
}

func (f FriendshipController) AcceptFriendshipRequest(_ context.Context, input *pb.AcceptFriendshipInput) (*pb.Friendship, error) {
	accept, err := f.friendshipService.AcceptFriendRequest(input.FriendshipId)
	if err != nil {
		return nil, err
	}
	return &pb.Friendship{
		Id: accept.ID.String(),
		User: &pb.User{
			Id:       accept.User.ID.String(),
			Username: accept.User.Username,
			Email:    accept.User.Email,
		},
		Friend: &pb.User{
			Id:       accept.Friend.ID.String(),
			Username: accept.Friend.Username,
			Email:    accept.Friend.Email,
		},
		Status: uint32(accept.Status),
	}, nil
}

func (f FriendshipController) RemoveFriendshipRequest(_ context.Context, input *pb.RemoveFriendshipInput) (*pb.Empty, error) {
	err := f.friendshipService.RemoveFriend(input.FriendshipId)
	if err != nil {
		log.Printf("Error removing friendship: %v", err)
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (f FriendshipController) FriendshipRequestList(_ context.Context, filter *pb.FriendshipFilter) (*pb.PaginatedFriendshipList, error) {
	friendships, err := f.friendshipService.GetFriendships(filter.UserId, ent.FriendshipStatus(filter.Status))
	if err != nil {
		return nil, err
	}

	var friendshipList []*pb.Friendship
	for _, friendship := range friendships {
		friendshipList = append(friendshipList, &pb.Friendship{
			Id: friendship.ID.String(),
			User: &pb.User{
				Id:       friendship.User.ID.String(),
				Username: friendship.User.Username,
				Email:    friendship.User.Email,
			},
			Friend: &pb.User{
				Id:       friendship.Friend.ID.String(),
				Username: friendship.Friend.Username,
				Email:    friendship.Friend.Email,
			},
			Status: uint32(friendship.Status),
		})
	}

	return &pb.PaginatedFriendshipList{
		Friendships: friendshipList,
		PaginationInfo: &pb.PaginationInfo{
			Count: int32(len(friendshipList)),
		},
	}, nil
}
