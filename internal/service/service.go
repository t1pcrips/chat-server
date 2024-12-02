package service

import (
	"chat-server/internal/model"
	"context"
)

type ChatService interface {
	Create(context.Context, []model.User) (int64, error)
	Delete(ctx context.Context, chatId int64) error
	SendMessage(ctx context.Context, info *model.CreateMessageRequest) error
}
