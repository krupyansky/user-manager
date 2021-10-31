package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/krupyansky/user-manager/internal/dto"
	"github.com/krupyansky/user-manager/internal/queue"
	pb "github.com/krupyansky/user-manager/pkg"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"strconv"
)

func (h *Handler) GetUsers(_ context.Context, _ *empty.Empty) (*pb.GetUsersResponse, error) {
	users, err := h.services.Authorization.GetUsers()
	if err != nil {
		log.Fatalln(err)
	}

	var response pb.GetUsersResponse

	for _, user := range users {
		u := pb.User{
			Id:    int32(user.Id),
			Name:  user.Name,
			Email: user.Email,
		}
		response.Users = append(response.Users, &u)
	}

	return &response, nil
}

func (h *Handler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	command := dto.UserProfile{
		Name:  req.Name,
		Email: req.Email,
	}

	id, _ := h.services.Authorization.CreateUser(command)

	w := kafka.Writer{Topic: queue.Topic, Addr: kafka.TCP(queue.BrokerAddress)}

	err := w.WriteMessages(ctx, kafka.Message{
		Key:   []byte(strconv.Itoa(id)),
		Value: []byte("this is message:" + req.Name + " " + req.Email),
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}

	return &pb.CreateUserResponse{Id: int32(id)}, nil
}

func (h *Handler) DeleteUser(_ context.Context, req *pb.DeleteUserRequest) (*empty.Empty, error) {
	command := dto.UserId{Id: int(req.Id)}

	err := h.services.Authorization.DeleteUser(command)
	if err != nil {
		log.Fatalln(err)
	}

	return new(emptypb.Empty), nil
}
