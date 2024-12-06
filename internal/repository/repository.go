package repository

import (
	"context"
	"github.com/t1pcrips/chat-service/internal/model"
)

type ChatRepository interface {
	CreateChat(ctx context.Context) (int64, error)
	DeleteChat(ctx context.Context, chatId int64) error
	CheckChatExists(ctx context.Context, chatId int64) (bool, error)
}

type MembersRepository interface {
	CreateMember(ctx context.Context, info *model.CreateMemberRequest) error
	CheckMemberExists(ctx context.Context, userName string) (bool, error)
}

type MessageRepository interface {
	CreateMessage(ctx context.Context, info *model.CreateMessageRequest) error
}
