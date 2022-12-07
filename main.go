package main

import (
	"log"
	"net"

	"github.com/Hwisaek/go-grpc/chat"
	"github.com/Hwisaek/go-grpc/user"
	"google.golang.org/grpc"
)

const portNumber = "19000"

func main() {
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	chat.RegisterChatServer(grpcServer, &chat.Service{})
	user.RegisterUserServer(grpcServer, &user.Service{})

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
