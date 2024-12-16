package chat

import (
	"context"
	"github.com/t1pcrips/chat-service/internal/converter"
	desc "github.com/t1pcrips/chat-service/pkg/chat_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *ChatApiImpl) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	err := i.service.SendMessage(ctx, converter.ToMessageFromSendApi(req))
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &emptypb.Empty{}, nil
}
