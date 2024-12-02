package ctrl

import (
	pb "app/internal/core/grpc/generated"
	"app/internal/pkg/friendship/svc"
	"context"
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
	//TODO implement me
	//panic("implement me")
	return nil, nil
}

func (f FriendshipController) RemoveFriendshipRequest(_ context.Context, input *pb.RemoveFriendshipInput) (*pb.Empty, error) {
	//TODO implement me
	//panic("implement me")
	return nil, nil
}

func (f FriendshipController) FriendshipRequestList(_ context.Context, filter *pb.FriendshipFilter) (*pb.PaginatedFriendshipList, error) {
	//TODO implement me
	//panic("implement me")
	return nil, nil
}
