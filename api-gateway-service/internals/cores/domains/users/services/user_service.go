package services

import (
	"context"
	"time"

	users "github.com/lakeside763/synapse-crm-pkg/protos/users"
	"google.golang.org/grpc"
	"github.com/lakeside763/synapse-crm-pkg/models"
)

type UserServiceClient struct {
	client users.UserServiceClient
}

func NewUserServiceClient(address string) (*UserServiceClient, error) {
	// Setup a connection to the UserService gRPC server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	
	client := users.NewUserServiceClient(conn)
	
	return &UserServiceClient{client},nil
}

func (u *UserServiceClient) CreateUser(data *models.CreateUserInput) (*users.CreateUserResponse, error) {
	// Create a new user via gRPC client
	req := &users.CreateUserRequest{
		Name: data.Name,
		Email: data.Email,
		Role: data.Role,
		Status: data.Status,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	return u.client.CreateUser(ctx, req)
}