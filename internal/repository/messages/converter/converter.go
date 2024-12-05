package converter

import (
	"github.com/t1pcrips/chat-service/internal/model"
	"github.com/t1pcrips/chat-service/internal/repository/messages/message_model"
)

func ToMessageFromMessageService(message *model.CreateMessageRequest) *message_model.CreateMessageRequest {
	return &message_model.CreateMessageRequest{
		ChatId:    message.ChatId,
		From:      message.From,
		Text:      message.Text,
		Timestamp: message.Timestamp,
	}
}
