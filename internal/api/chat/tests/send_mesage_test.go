package tests

import (
	"chat-server/internal/api/chat"
	"chat-server/internal/model"
	"chat-server/internal/service"
	"chat-server/internal/service/mocks"
	desc "chat-server/pkg/chat_v1"
	"chat-server/pkg/errs"
	"context"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math/rand"
	"testing"
	"time"
)

func TestSendMessage(t *testing.T) {
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *desc.SendMessageRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		from      = "Timofey"
		text      = "Hello my friend wowowowo"
		timestamp = timestamppb.New(time.Now())
		chatId    = rand.Int63()

		req = &desc.SendMessageRequest{
			From:      from,
			Text:      text,
			Timestamp: timestamp,
			ToChatId:  chatId,
		}

		info = &model.CreateMessageRequest{
			ChatId:    chatId,
			From:      from,
			Text:      text,
			Timestamp: timestamp.AsTime(),
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: successSend,
			args: args{
				ctx: ctx,
				req: req,
			},
			want: &emptypb.Empty{},
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := mocks.NewChatServiceMock(mc)
				mock.SendMessageMock.Expect(ctx, info).Return(nil)
				return mock
			},
		},
		{
			name: failedSend,
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  status.Error(codes.NotFound, errs.ErrFailedService.Error()),
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := mocks.NewChatServiceMock(mc)
				mock.SendMessageMock.Expect(ctx, info).Return(errs.ErrFailedService)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatServiceMock := tt.chatServiceMock(mc)
			api := chat.NewChatApiImpl(chatServiceMock)

			testResponse, err := api.SendMessage(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.want, testResponse)
			require.Equal(t, tt.err, err)
		})
	}
}