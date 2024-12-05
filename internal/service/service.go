package service

import (
	"context"
	"github.com/t1pcrips/chat-service/internal/model"
)

type ChatService interface {
	Create(ctx context.Context, info *model.CreateChatRequest) (int64, error)
	Delete(ctx context.Context, chatId int64) error
	SendMessage(ctx context.Context, info *model.CreateMessageRequest) error
}
