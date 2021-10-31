package main

import (
	"context"
	"fmt"
	pb "github.com/krupyansky/user-manager/pkg"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	client := pb.NewUserApiClient(conn)

	resp, err := client.CreateUser(context.Background(), &pb.CreateUserRequest{Name: "Slava", Email: "mail2@mail.ru"})
	if err != nil {
		return
	}

	fmt.Println(resp)
	fmt.Println(resp.Id)
}
