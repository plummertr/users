package users

import (
	"context"

	v1 "github.com/plummertr/users/api/internal-api/v1"
)

type userServer struct {
	v1.UnimplementedUserServiceServer
}

func NewUserServer() *userServer {
	return &userServer{}
}

func (u userServer) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	res := v1.CreateUserResponse{
		Id: int64(1),
	}

	return &res, nil
}

func (u userServer) ReadUser(ctx context.Context, req *v1.ReadUserRequest) (*v1.ReadUserResponse, error) {
	user := v1.User{
		Id:       1,
		Name:     "Tom Plummer",
		Email:    "plummertr@gmail.com",
		Username: "plummert",
	}

	res := v1.ReadUserResponse{
		User: &user,
	}

	return &res, nil
}

func (u userServer) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {
	return &v1.UpdateUserResponse{}, nil
}

func (u userServer) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserResponse, error) {
	return &v1.DeleteUserResponse{}, nil
}
