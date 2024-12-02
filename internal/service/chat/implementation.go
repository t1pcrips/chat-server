package chat

import (
	"chat-server/internal/client/database"
	"chat-server/internal/repository"
	"chat-server/internal/service"
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
