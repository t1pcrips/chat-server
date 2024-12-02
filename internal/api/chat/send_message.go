package chat

import (
	"chat-server/internal/converter"
	dst "chat-server/pkg/chat_v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *ChatApiImpl) SendMessage(ctx context.Context, req *dst.SendMessageRequest) (*emptypb.Empty, error) {
	err := i.service.SendMessage(ctx, converter.ToMessageFromDeps(req))
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &emptypb.Empty{}, nil
}
