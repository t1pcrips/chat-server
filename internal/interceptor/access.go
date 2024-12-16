package interceptor

import (
	"context"
	"errors"
	"github.com/t1pcrips/chat-service/internal/client"
	"github.com/t1pcrips/chat-service/internal/utils"
	"google.golang.org/grpc"
)

type AccessInterceptor struct {
	client client.AccessClient
}

func NewAccessInterceptor(client client.AccessClient) *AccessInterceptor {
	return &AccessInterceptor{
		client: client,
	}
}

func (i *AccessInterceptor) Check(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	outgoingCtx, err := utils.FromIncomingToOutgoingCtx(ctx)
	if err != nil {
		return nil, err
	}

	err = i.client.Check(outgoingCtx, info.FullMethod)
	if err != nil {
		return nil, errors.New("access denied")
	}

	return handler(ctx, req)
}
