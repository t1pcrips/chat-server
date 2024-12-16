package converter

import (
	"context"
	"github.com/t1pcrips/chat-service/internal/model"
	"github.com/t1pcrips/chat-service/internal/service"
	"github.com/t1pcrips/chat-service/pkg/chat_v1"
)

type StreamChatMessages struct {
	chatServer chat_v1.Chat_ConnectServer
}

func NewStreamChatMessages(chatService chat_v1.Chat_ConnectServer) service.StreamChatMessages {
	return &StreamChatMessages{
		chatServer: chatService,
	}
}

func (c *StreamChatMessages) Send(message *model.Message) error {
	return c.chatServer.Send(ToApiFromMessageSender(message))
}

func (c *StreamChatMessages) Context() context.Context {
	return c.chatServer.Context()
}
