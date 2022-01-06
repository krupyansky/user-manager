package main

import (
	"context"
	"fmt"
	pb "github.com/krupyansky/user-manager/pkg"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	client := pb.NewUserApiClient(conn)

	resp, err := client.CreateUser(context.Background(), &pb.CreateUserRequest{Name: "Suzana", Email: "fixin@mail.ru"})
	if err != nil {
		return
	}

	fmt.Println(resp)
	fmt.Println(resp.Id)

	//resp, err := client.DeleteUser(context.Background(), &pb.DeleteUserRequest{Id: 5})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(resp)

	//resp, err := client.GetUsers(context.Background(), &empty.Empty{})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(resp.Users)
}
