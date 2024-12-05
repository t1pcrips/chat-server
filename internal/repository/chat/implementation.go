package chat

import (
	"github.com/t1pcrips/chat-service/internal/repository"
	"github.com/t1pcrips/platform-pkg/pkg/database"
)

const (
	tableChats      = "chats"
	idColumn        = "id"
	createdAtColumn = "created_at"
	returningId     = "RETURNING id"
)

type ChatRepositoryImpl struct {
	db database.Client
}

var _ repository.ChatRepository = (*ChatRepositoryImpl)(nil)

func NewChatRepositoryImpl(db database.Client) repository.ChatRepository {
	return &ChatRepositoryImpl{
		db: db,
	}
}
