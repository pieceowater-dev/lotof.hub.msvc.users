package ctrl

import (
	pb "app/internal/core/grpc/generated"
	"app/internal/pkg/user/ent"
	"app/internal/pkg/user/svc"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"net/http"
)

type UserController struct {
	userService *svc.UserService
	pb.UnimplementedUserServiceServer
}

func NewUserController(service *svc.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

func (c UserController) GetUsers(_ context.Context, request *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	filter := gossiper.NewFilter[string](
		request.GetSearch(),
		gossiper.NewSort[string](
			request.GetSort().GetField(),
			gossiper.SortDirection(request.GetSort().GetDirection()),
		),
		gossiper.NewPagination(
			int(request.GetPagination().GetPage()),
			int(request.GetPagination().GetLength()),
		),
	)

	paginatedResult, err := c.userService.GetUsers(filter)
	if err != nil {
		return nil, err
	}

	return &pb.GetUsersResponse{
		Users: paginatedResult.Rows,
		PaginationInfo: &pb.PaginationInfo{
			Count: int32(paginatedResult.Info.Count),
		},
	}, nil
}

func (c UserController) GetUser(_ context.Context, request *pb.GetUserRequest) (*pb.User, error) {
	user, err := c.userService.GetUserByID(request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Id:       user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (c UserController) UpdateUser(_ context.Context, request *pb.UpdateUserRequest) (*pb.User, error) {
	user, err := c.userService.UpdateUser(&ent.User{
		ID:       uuid.MustParse(request.Id),
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Id:       user.ID.String(),
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

// ListREST handles the HTTP request for getting users via Gin.
func (c UserController) ListREST(ctx *gin.Context) {
	res, err := c.GetUsers(ctx, &pb.GetUsersRequest{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(200, res)
}
