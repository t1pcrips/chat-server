package chat

import (
	dst "chat-server/pkg/chat_v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *ChatApiImpl) Delete(ctx context.Context, req *dst.DeleteRequest) (*emptypb.Empty, error) {
	err := i.service.Delete(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &emptypb.Empty{}, nil
}
