package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/lakeside763/synapse-crm-pkg/models"
	users "github.com/lakeside763/synapse-crm-pkg/protos/users"
	"github.com/lakeside763/synapse-crm/user-service/internals/ports/interfaces"
)

// Correctly embedding UnimplementedUserServiceServer
type UserServiceServer struct {
	users.UnimplementedUserServiceServer
	repo interfaces.UserInterface
}

// Constructor function for UserService
func NewUserServiceServer(repo interfaces.UserInterface) *UserServiceServer {
	return &UserServiceServer{
		repo: repo,
	}
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *users.CreateUserRequest) (*users.CreateUserResponse, error) {
	// Do some logic here
	user := &models.User{
		ID:     uuid.New(),
		Name:   req.Name,
		Email:  req.Email,
		Status: req.Status,
		Role:   req.Role,
	}

	return &users.CreateUserResponse{
		Id: user.ID.String(),
		Name: user.Name,
		Email: user.Email,
		Status: user.Status,
		Role: user.Role,
	}, nil
}
