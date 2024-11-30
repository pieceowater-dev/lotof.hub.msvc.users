package ctrl

import (
	"app/internal/pkg/user/ent"
	"app/internal/pkg/user/svc"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *svc.UserService
}

func NewUserController(service *svc.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

// GetUsers retrieves the list of users (abstracted from any HTTP framework).
func (c UserController) GetUsers() []ent.User {
	// Fetch users from database or any other source
	users := []ent.User{}
	// Add logic to fetch users, for example:
	//users, err := c.userService.GetAllUsers()
	//if err != nil {
	//   return nil
	//}
	return users
}

// ListREST handles the HTTP request for getting users via Gin.
func (c UserController) ListREST(ctx *gin.Context) {
	ctx.JSON(200, c.GetUsers())
}
