package utils

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
)

func FromIncomingToOutgoingCtx(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata access inceptor")
	}

	return metadata.NewOutgoingContext(ctx, md), nil
}
