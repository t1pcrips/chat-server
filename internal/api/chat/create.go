package chat

import (
	"chat-server/internal/converter"
	dst "chat-server/pkg/chat_v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *ChatApiImpl) Create(ctx context.Context, req *dst.CreateRequest) (*dst.CreateResponse, error) {
	chatId, err := i.service.Create(ctx, converter.ConvertUsernamesToUsers(req.GetUsernames()))
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &dst.CreateResponse{
		Id: chatId,
	}, nil
}
