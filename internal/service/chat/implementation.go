package chat

import (
	"github.com/t1pcrips/chat-service/internal/repository"
	"github.com/t1pcrips/chat-service/internal/service"
	"github.com/t1pcrips/platform-pkg/pkg/database"
)

type ChatServiceImpl struct {
	chatRepository    repository.ChatRepository
	memberRepository  repository.MembersRepository
	messageRepository repository.MessageRepository
	txManeger         database.TxManeger
}

func NewChatService(
	chatRepository repository.ChatRepository,
	memberRepository repository.MembersRepository,
	messageRepository repository.MessageRepository,
	txManeger database.TxManeger,
) service.ChatService {
	return &ChatServiceImpl{
		chatRepository:    chatRepository,
		memberRepository:  memberRepository,
		messageRepository: messageRepository,
		txManeger:         txManeger,
	}
}
