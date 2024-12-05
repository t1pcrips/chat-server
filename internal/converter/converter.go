package converter

import (
	"github.com/t1pcrips/chat-service/internal/model"
	desc "github.com/t1pcrips/chat-service/pkg/chat_v1"
)

func ConvertToCreateRequestFromDesc(usernames []string) *model.CreateChatRequest {
	users := make([]string, len(usernames))
	for i, username := range usernames {
		users[i] = username
	}
	return &model.CreateChatRequest{
		Usernames: usernames,
	}
}

func ToMessageFromDeps(req *desc.SendMessageRequest) *model.CreateMessageRequest {
	return &model.CreateMessageRequest{
		ChatId:    req.GetToChatId(),
		From:      req.GetFrom(),
		Text:      req.GetText(),
		Timestamp: req.GetTimestamp().AsTime(),
	}
}
