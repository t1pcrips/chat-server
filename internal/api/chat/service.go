package chat

import (
	"github.com/t1pcrips/chat-service/internal/service"
	"github.com/t1pcrips/chat-service/pkg/chat_v1"
)

type ChatApiImpl struct {
	chat_v1.UnimplementedChatServer
	service service.ChatService
}

func NewChatApiImpl(service service.ChatService) *ChatApiImpl {
	return &ChatApiImpl{
		service: service,
	}
}
