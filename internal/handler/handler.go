package handler

import (
	_ "github.com/krupyansky/todo-app/docs"
	"github.com/krupyansky/user-manager/internal/service"
	pb "github.com/krupyansky/user-manager/pkg"
)

type Handler struct {
	pb.UnimplementedUserApiServer
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
