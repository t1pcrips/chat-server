package chat

import (
	"chat-server/internal/converter"
	desc "chat-server/pkg/chat_v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *ChatApiImpl) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	chatId, err := i.service.Create(ctx, converter.ConvertToCreateRequestFromDesc(req.GetUsernames()))
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &desc.CreateResponse{
		Id: chatId,
	}, nil
}
