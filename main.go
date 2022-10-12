package main

import (
	"context"
	"github.com/Hwisaek/go-grpc/data"
	"github.com/Hwisaek/go-grpc/message"
	"log"
	"net"

	"google.golang.org/grpc"
)

const portNumber = "9000"

type messageServer struct {
	message.UserServer
}

func (s *messageServer) SendMessage(ctx context.Context, req *message.SendMessageRequest) (*message.SendMessageResponse, error) {
	data.Message = append(data.Message, req.Message)

	return &message.SendMessageResponse{}, nil
}

func (s *messageServer) GetMessage(ctx context.Context, req *message.GetMessagesRequest) (*message.GetMessagesResponse, error) {
	messages := make([]*message.Message, len(data.Message))
	for i, m := range data.Message {
		messages[i] = m
	}

	return &message.GetMessagesResponse{Messages: messages}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	message.RegisterUserServer(grpcServer, &messageServer{})

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
