package presenter

import (
	"context"
	"fmt"
	proto "github.com/robertgarayshin/driveshare/proto/users"
)

type UserServer struct {
	proto.UsersServer
}

func (s UserServer) Auth(ctx context.Context, user *proto.User) (*proto.Token, error) {
	return nil, nil
}

func (s UserServer) Create(ctx context.Context, user *proto.User) (*proto.UserResponse, error) {
	fmt.Println(user)
	return &proto.UserResponse{User: user}, nil
}

func (s UserServer) Get(ctx context.Context, user *proto.User) (*proto.UserResponse, error) {
	fmt.Println("test request")
	if user.Id == 1 {
		return &proto.UserResponse{User: &proto.User{Id: 1, Email: "test@test.com"}}, nil
	}
	return nil, nil
}

func (s UserServer) GetAll(ctx context.Context, req *proto.Request) (*proto.UserResponse, error) {
	return nil, nil
}

func (s UserServer) ValidateToken(ctx context.Context, token *proto.Token) (*proto.Token, error) {
	return nil, nil
}
