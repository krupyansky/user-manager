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

	//resp, err := client.CreateUser(context.Background(), &pb.CreateUserRequest{Name: "Slava", Email: "mail87@mail.ru"})
	//if err != nil {
	//	return
	//}
	//
	//fmt.Println(resp)
	//fmt.Println(resp.Id)

	//resp, err := client.DeleteUser(context.Background(), &pb.DeleteUserRequest{Id: 5})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(resp)

	resp, err := client.GetUsers(context.Background(), &empty.Empty{})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.Users)
}
