package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
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

	resp, err := client.GetUsers(context.Background(), &empty.Empty{})
	if err != nil {
		return
	}

	fmt.Println(resp)
	fmt.Println(resp.Users[0].Email)
}
