package converter

import (
	"github.com/t1pcrips/chat-service/internal/model"
	"github.com/t1pcrips/chat-service/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func ToCreateRequestFromApi(usernames []string) *model.CreateChatRequest {
	users := make([]string, len(usernames))
	for i, username := range usernames {
		users[i] = username
	}
	return &model.CreateChatRequest{
		Usernames: usernames,
	}
}

func ToMessageFromSendApi(info *chat_v1.SendMessageRequest) *model.CreateMessageRequest {
	return &model.CreateMessageRequest{
		ChatId:    info.GetToChatId(),
		From:      info.GetFrom(),
		Text:      info.GetText(),
		Timestamp: info.GetTimestamp().AsTime(),
	}
}

func ToUserFromConnectApi(info *chat_v1.ConnectChatRequest) *model.User {
	return &model.User{
		ChatId:   info.GetChatId(),
		Username: info.GetUsername(),
	}
}

func ToApiFromMessageSender(info *model.Message) *chat_v1.Message {
	return &chat_v1.Message{
		ToChatId:  info.ChatId,
		From:      info.From,
		Text:      info.Text,
		Timestamp: timestamppb.New(time.Now()),
	}
}
