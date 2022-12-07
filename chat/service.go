package chat

import (
	"context"

	"github.com/Hwisaek/go-grpc/data"
)

type Service struct {
	ChatServer
}

func (s *Service) SendMessage(ctx context.Context, req *SendMessageRequest) (res *SendMessageResponse, err error) {
	msg := data.Message{
		Name: req.Message.Name,
		Text: req.Message.Text,
	}
	data.MessageList = append(data.MessageList, msg)

	return &SendMessageResponse{}, nil
}

func (s *Service) GetMessage(ctx context.Context, req *GetMessagesRequest) (res *GetMessagesResponse, err error) {
	messages := make([]*Message, len(data.MessageList))
	for i, m := range data.MessageList {
		messages[i] = &Message{
			Name: m.Name,
			Text: m.Text,
		}
	}

	return &GetMessagesResponse{Messages: messages}, nil
}
