package chat

import (
	"chat-server/internal/service"
	dst "chat-server/pkg/chat_v1"
)

type ChatApiImpl struct {
	dst.UnimplementedChatServer
	service service.ChatService
}

func NewChatApiImpl(service service.ChatService) *ChatApiImpl {
	return &ChatApiImpl{
		service: service,
	}
}
