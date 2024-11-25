package ctrl

import "app/internal/pkg/friendship/svc"

type FriendshipController struct {
	friendshipService *svc.FriendshipService
}

func NewFriendshipController(service *svc.FriendshipService) *FriendshipController {
	return &FriendshipController{
		friendshipService: service,
	}
}

// todo: implement methods
