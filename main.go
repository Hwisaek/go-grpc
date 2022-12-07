package main

import (
	"context"
	"errors"
	"log"
	"net"

	"github.com/Hwisaek/go-grpc/chat"
	"github.com/Hwisaek/go-grpc/data"
	"github.com/Hwisaek/go-grpc/user"
	"google.golang.org/grpc"
)

const portNumber = "19000"

type chatServer struct {
	chat.ChatServer
}

func (s *chatServer) SendMessage(ctx context.Context, req *chat.SendMessageRequest) (*chat.SendMessageResponse, error) {
	data.Message = append(data.Message, req.Message)

	return &chat.SendMessageResponse{}, nil
}

func (s *chatServer) GetMessage(ctx context.Context, req *chat.GetMessagesRequest) (*chat.GetMessagesResponse, error) {
	messages := make([]*chat.Message, len(data.Message))
	for i, m := range data.Message {
		messages[i] = m
	}

	return &chat.GetMessagesResponse{Messages: messages}, nil
}

type userServer struct {
	user.UserServer
}

func (s *userServer) PostLogin(ctx context.Context, req *user.PostLoginRequest) (*user.PostLoginResponse, error) {
	if req.Id != "test" || req.Pw != "1234" {
		return nil, errors.New("login error")
	}
	return &user.PostLoginResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	chat.RegisterChatServer(grpcServer, &chatServer{})
	user.RegisterUserServer(grpcServer, &userServer{})

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
