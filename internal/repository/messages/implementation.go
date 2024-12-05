package messages

import (
	"github.com/t1pcrips/chat-service/internal/repository"
	"github.com/t1pcrips/platform-pkg/pkg/database"
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
