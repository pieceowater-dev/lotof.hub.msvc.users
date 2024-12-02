package svc

import (
	pb "app/internal/core/grpc/generated"
	"app/internal/pkg/user/ent"
	"errors"
	"fmt"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db gossiper.Database
}

func NewUserService(db gossiper.Database) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(user *ent.User) (*ent.User, error) { // todo: delete this later
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
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

func (s *UserService) GetUsers(filter gossiper.Filter[string]) (gossiper.PaginatedResult[*pb.User], error) {
	var users []ent.User
	var count int64

	query := s.db.GetDB().Model(&ent.User{})

	// Apply search filters
	if filter.Search != "" {
		search := "%" + filter.Search + "%"
		query = query.Where("username LIKE ? OR email LIKE ?", search, search)
	}

	// Count total records
	if err := query.Count(&count).Error; err != nil {
		return gossiper.PaginatedResult[*pb.User]{}, fmt.Errorf("failed to count users: %w", err)
	}

	// Apply pagination
	query = query.Offset((filter.Pagination.Page - 1) * filter.Pagination.Length).Limit(filter.Pagination.Length)

	// Apply sorting dynamically
	if field := filter.Sort.Field; field != "" && gossiper.IsFieldValid(&ent.User{}, field) {
		query = query.Order(fmt.Sprintf("%s %s", gossiper.ToSnakeCase(field), filter.Sort.Direction))
	}

	// Fetch data
	if err := query.Find(&users).Error; err != nil {
		return gossiper.PaginatedResult[*pb.User]{}, fmt.Errorf("failed to fetch users: %w", err)
	}

	// Map results
	grpcUsers := make([]*pb.User, len(users))
	for i, u := range users {
		grpcUsers[i] = &pb.User{
			Id:       u.ID.String(),
			Username: u.Username,
			Email:    u.Email,
		}
	}

	// Create paginated result
	return gossiper.NewPaginatedResult(grpcUsers, int(count)), nil
}
