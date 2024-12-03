package chat

import (
	"chat-server/internal/service"
	desc "chat-server/pkg/chat_v1"
)

type ChatApiImpl struct {
	desc.UnimplementedChatServer
	service service.ChatService
}

func NewChatApiImpl(service service.ChatService) *ChatApiImpl {
	return &ChatApiImpl{
		service: service,
	}
}
