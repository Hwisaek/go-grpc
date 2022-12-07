package chat

import (
	"context"
	"time"

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

func (s *Service) GetMessageStream(req *GetMessagesRequest, stream Chat_GetMessageStreamServer) error {
	messages := make([]*Message, len(data.MessageList))
	for i, m := range data.MessageList {
		messages[i] = &Message{
			Name: m.Name,
			Text: m.Text,
		}
	}
	i := 0
	for {
		if i == 10 {
			break
		}
		stream.Send(&GetMessagesResponse{Messages: messages})
		time.Sleep(time.Second / 10)
		i++
	}
	return nil
}
