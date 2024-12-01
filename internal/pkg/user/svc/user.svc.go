package svc

import (
	pb "app/internal/core/grpc/generated"
	"app/internal/pkg/user/ent"
	"errors"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"gorm.io/gorm"
)

type UserService struct {
	db gossiper.Database
}

func NewUserService(db gossiper.Database) *UserService {
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

// todo: make correct filter instead of pb.GetUsersRequest

func (s *UserService) GetUsers(request *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	var users []ent.User
	var count int64

	query := s.db.GetDB().Model(&ent.User{})

	if request.Search != "" {
		search := "%" + request.Search + "%"
		query = query.Where("username LIKE ? OR email LIKE ?", search, search)
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, err
	}

	offset := int(request.Pagination.Page-1) * int(request.Pagination.Length)
	query = query.Offset(offset).Limit(int(request.Pagination.Length))

	if request.Sort != nil && request.Sort.Field != "" {
		sortDirection := "ASC"
		if request.Sort.Direction == "DESC" {
			sortDirection = "DESC"
		}
		query = query.Order(request.Sort.Field + " " + sortDirection)
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}

	var grpcUsers []*pb.User
	for _, u := range users {
		grpcUsers = append(grpcUsers, &pb.User{
			Id:       u.ID.String(),
			Username: u.Username,
			Email:    u.Email,
		})
	}

	return &pb.GetUsersResponse{
		Users: grpcUsers,
		PaginationInfo: &pb.PaginationInfo{
			Count: int32(count),
		},
	}, nil
}

//func (s *UserService) AddFriend(userID, friendID string) error {
//	var user, friend ent.User
//	if err := s.db.GetDB().First(&user, "id = ?", userID).Error; err != nil {
//		return err
//	}
//	if err := s.db.GetDB().First(&friend, "id = ?", friendID).Error; err != nil {
//		return err
//	}
//	return s.db.GetDB().Model(&user).Association("Friends").Append(&friend)
//}
