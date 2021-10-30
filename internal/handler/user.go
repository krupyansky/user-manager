package handler

import (
	"context"
	pb "github.com/krupyansky/user-manager/pkg"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Handler) GetUsers(context.Context, *emptypb.Empty) (*pb.GetUsersResponse, error) {
	var response pb.GetUsersResponse

	var user pb.User
	user.Id = 1
	user.Name = "Slava"
	user.Email = "mail@mail.ru"

	users := append(response.Users, &user)

	response = pb.GetUsersResponse{Users: users}

	return &response, nil
}

func (h *Handler) CreateUser(context.Context, *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	var response pb.CreateUserResponse
	return &response, nil
}

func (h *Handler) DeleteUser(context.Context, *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	return nil, nil
}
