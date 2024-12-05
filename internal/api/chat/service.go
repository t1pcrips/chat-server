package chat

import (
	"github.com/t1pcrips/chat-service/internal/service"
	desc "github.com/t1pcrips/chat-service/pkg/chat_v1"
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
