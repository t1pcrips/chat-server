package tests

/*

import (
	"github.com/t1pcrips/chat-service/internal/model"
	"github.com/t1pcrips/chat-service/internal/repository"
	"github.com/t1pcrips/chat-service/internal/repository/mocks"
	"github.com/t1pcrips/chat-service/internal/service/chat"
	"github.com/t1pcrips/chat-service/pkg/errs"
	"github.com/t1pcrips/chat-service/pkg/helpers"
	"context"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

func TestSendMessage(t *testing.T) {
	type sendMessageServiceMockFunc func(mc *minimock.Controller) repository.MessageRepository

	type args struct {
		ctx context.Context
		req *model.CreateMessageRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		chatId    = rand.Int63()
		from      = "Timofey"
		text      = "Hello my friend wowowowo"
		timestamp = time.Now()

		req = &model.CreateMessageRequest{
			ChatId:    chatId,
			From:      from,
			Text:      text,
			Timestamp: timestamp,
		}

		info = &model.CreateMessageRequest{
			ChatId:    chatId,
			From:      from,
			Text:      text,
			Timestamp: timestamp,
		}
	)

	tests := []struct {
		name                   string
		args                   args
		err                    error
		sendMessageServiceMock sendMessageServiceMockFunc
	}{
		{
			name: helpers.SuccessSend,
			args: args{
				ctx: ctx,
				req: req,
			},
			err: nil,
			sendMessageServiceMock: func(mc *minimock.Controller) repository.MessageRepository {
				mock := mocks.NewMessageRepositoryMock(mc)
				mock.CreateMessageMock.Expect(ctx, info).Return(nil)

				return mock
			},
		},
		{
			name: helpers.FailedSend,
			args: args{
				ctx: ctx,
				req: req,
			},
			err: errs.ErrSendMessage,
			sendMessageServiceMock: func(mc *minimock.Controller) repository.MessageRepository {
				mock := mocks.NewMessageRepositoryMock(mc)
				mock.CreateMessageMock.Expect(ctx, info).Return(errs.ErrSendMessage)

				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			sendMessageServiceMock := tt.sendMessageServiceMock(mc)

			service := chat.NewChatService(nil, nil, sendMessageServiceMock, nil)

			err := service.SendMessage(ctx, info)
			require.Equal(t, tt.err, err)
		})
	}
}
*/
