package user

import (
	"context"
	"errors"
)

type Service struct {
	UserServer
}

func (s *Service) PostLogin(ctx context.Context, req *PostLoginRequest) (*PostLoginResponse, error) {
	if req.Id != "test" || req.Pw != "1234" {
		return nil, errors.New("login error")
	}
	return &PostLoginResponse{}, nil
}
