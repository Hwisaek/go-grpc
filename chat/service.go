package chat

import (
	"context"

	"github.com/Hwisaek/go-grpc/data"
)

type Service struct {
	ChatServer
}

func (s *Service) SendMessage(ctx context.Context, req *SendMessageRequest) (*SendMessageResponse, error) {
	data.Message = append(data.Message, req.Message)

	return &SendMessageResponse{}, nil
}

func (s *Service) GetMessage(ctx context.Context, req *GetMessagesRequest) (*GetMessagesResponse, error) {
	messages := make([]*Message, len(data.Message))
	for i, m := range data.Message {
		messages[i] = m
	}

	return &GetMessagesResponse{Messages: messages}, nil
}
