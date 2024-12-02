package messages

import (
	"chat-server/internal/client/database"
	"chat-server/internal/repository"
)

const (
	tableMessages   = "messages"
	idColumn        = "id"
	chatIdColumn    = "chat_id"
	fromColumn      = "from_username"
	textColumn      = "text"
	timestampColumn = "timestamp"
)

type MessagesRepositoryImpl struct {
	db database.Client
}

func NewMessagesRepository(db database.Client) repository.MessageRepository {
	return &MessagesRepositoryImpl{
		db: db,
	}
}
