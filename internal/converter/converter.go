package converter

import (
	"chat-server/internal/model"
	dst "chat-server/pkg/chat_v1"
)

func ConvertUsernamesToUsers(usernames []string) []model.User {
	users := make([]model.User, len(usernames))
	for i, username := range usernames {
		users[i] = model.User{Username: username}
	}
	return users
}

func ToMessageFromDeps(req *dst.SendMessageRequest) *model.CreateMessageRequest {
	return &model.CreateMessageRequest{
		ChatId:    req.GetToChatId(),
		From:      req.GetFrom(),
		Text:      req.GetText(),
		Timestamp: req.GetTimestamp().AsTime(),
	}
}
