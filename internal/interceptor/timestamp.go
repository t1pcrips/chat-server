package interceptor

import (
	"context"
	desc "github.com/t1pcrips/chat-service/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func TimestampInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if val, ok := req.(*desc.SendMessageRequest); ok {
		now := time.Now()

		val.Timestamp = timestamppb.New(now)
	}

	return handler(ctx, req)
}
