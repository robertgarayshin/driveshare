package grpc

import (
	"context"
	pb "driveshare/pkg/model/proto/user"
	"user/usecase"
)

type UserHandler struct {
	usecase *usecase.UserUsecase
	pb.UnimplementedUserServiceServer
}

func NewUserHandler(u *usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: u}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user, err := h.usecase.CreateUser(req.Name, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := h.usecase.GetUser(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
