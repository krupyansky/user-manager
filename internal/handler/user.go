package handler

import (
	"context"
	"github.com/krupyansky/user-manager/internal/dto"
	pb "github.com/krupyansky/user-manager/pkg"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Handler) GetUsers(ctx context.Context, req *emptypb.Empty) (*pb.GetUsersResponse, error) {
	var response pb.GetUsersResponse

	var user pb.User
	user.Id = 1
	user.Name = "Slava"
	user.Email = "mail@mail.ru"

	users := append(response.Users, &user)

	response = pb.GetUsersResponse{Users: users}

	return &response, nil
}

func (h *Handler) CreateUser(_ context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	command := dto.UserProfile{
		Name:  req.Name,
		Email: req.Email,
	}

	id, _ := h.services.Authorization.CreateUser(command)

	return &pb.CreateUserResponse{Id: int32(id)}, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	return nil, nil
}
