package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"user-service"
)

func main() {
	list , err := net.Listen("tcp",":8080")
	if err != nil{
		fmt.Println(err)
	}
	server :=grpc.NewServer()
	user_service.RegisterUserServiceServer(server, &user_service.UserServiceServerImpl{})
	if err := server.Serve(list); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
