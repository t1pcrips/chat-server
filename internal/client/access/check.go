package access

import (
	"context"
	"github.com/t1pcrips/chat-service/internal/client"
	"github.com/t1pcrips/chat-service/pkg/access_v1"
)

type AccessClientImpl struct {
	client access_v1.AccessClient
}

func NewAccessClientImpl(client access_v1.AccessClient) client.AccessClient {
	return &AccessClientImpl{
		client: client,
	}
}

func (c *AccessClientImpl) Check(ctx context.Context, address string) error {
	_, err := c.client.Check(ctx, &access_v1.CheckRequest{
		Address: address,
	})
	if err != nil {
		return err
	}

	return nil
}
