package converter

import (
	"chat-server/internal/model"
	"chat-server/internal/repository/messages/message_model"
)

func ToMessageFromMessageService(message *model.CreateMessageRequest) *message_model.CreateMessageRequest {
	return &message_model.CreateMessageRequest{
		ChatId:    message.ChatId,
		From:      message.From,
		Text:      message.Text,
		Timestamp: message.Timestamp,
	}
}
