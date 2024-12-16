package chat

import (
	"github.com/t1pcrips/chat-service/internal/converter"
	"github.com/t1pcrips/chat-service/pkg/chat_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *ChatApiImpl) Connect(req *chat_v1.ConnectChatRequest, stream chat_v1.Chat_ConnectServer) error {
	err := i.service.Connect(converter.ToUserFromConnectApi(req), converter.NewStreamChatMessages(stream))
	if err != nil {
		return status.Error(codes.Internal, "some thing wrong")
	}

	return nil
}
