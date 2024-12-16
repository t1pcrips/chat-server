package chat

import (
	"github.com/t1pcrips/chat-service/internal/repository"
	"github.com/t1pcrips/chat-service/internal/service"
	"github.com/t1pcrips/chat-service/internal/service/chat/streams"
	"github.com/t1pcrips/platform-pkg/pkg/database"
)

type ChatServiceImpl struct {
	chatRepository    repository.ChatRepository
	memberRepository  repository.MembersRepository
	messageRepository repository.MessageRepository
	txManeger         database.TxManeger

	chats                *streams.Chats
	chatsMessageChannels *streams.ChatMessageChannels
}

func NewChatService(
	chatRepository repository.ChatRepository,
	memberRepository repository.MembersRepository,
	messageRepository repository.MessageRepository,
	txManeger database.TxManeger,
	chats *streams.Chats,
	chatsMessageChannels *streams.ChatMessageChannels,
) service.ChatService {
	return &ChatServiceImpl{
		chatRepository:       chatRepository,
		memberRepository:     memberRepository,
		messageRepository:    messageRepository,
		txManeger:            txManeger,
		chats:                chats,
		chatsMessageChannels: chatsMessageChannels,
	}
}
