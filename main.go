package main

import (
	"context"
	"errors"
	"log"
	"net"
	"time"

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

	grpcServer := grpc.NewServer(grpc.StreamInterceptor(func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		log.Print(info)
		return handler(srv, ss)
	}), grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		log.Print(info)
		if _, ok := req.(*user.PostLoginRequest); ok {
			return nil, errors.New("login error")
		}
		return handler(ctx, req)
	}), grpc.ConnectionTimeout(time.Nanosecond))
	chat.RegisterChatServer(grpcServer, &chat.Service{})
	user.RegisterUserServer(grpcServer, &user.Service{})

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
