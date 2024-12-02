package chat

import (
	"chat-server/internal/client/database"
	"chat-server/internal/repository"
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
