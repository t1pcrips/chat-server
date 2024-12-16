package chat

import (
	"context"
	"github.com/t1pcrips/chat-service/internal/converter"
	desc "github.com/t1pcrips/chat-service/pkg/chat_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *ChatApiImpl) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	chatId, err := i.service.Create(ctx, converter.ToCreateRequestFromApi(req.GetUsernames()))
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &desc.CreateResponse{
		Id: chatId,
	}, nil
}
