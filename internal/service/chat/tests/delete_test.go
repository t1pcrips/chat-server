package tests

import (
	"context"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"github.com/t1pcrips/chat-service/internal/repository"
	"github.com/t1pcrips/chat-service/internal/repository/mocks"
	"github.com/t1pcrips/chat-service/internal/service/chat"
	"github.com/t1pcrips/chat-service/pkg/errs"
	"github.com/t1pcrips/chat-service/pkg/helpers"
	"math/rand"
	"testing"
)

func TestDelete(t *testing.T) {
	type chatServiceMockFunc func(mc *minimock.Controller) repository.ChatRepository

	type args struct {
		ctx context.Context
		req int64
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		chatId = rand.Int63()
	)

	tests := []struct {
		name            string
		args            args
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: helpers.SuccessDelete,
			args: args{
				ctx: ctx,
				req: chatId,
			},
			err: nil,
			chatServiceMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := mocks.NewChatRepositoryMock(mc)
				mock.DeleteChatMock.Expect(ctx, chatId).Return(nil)

				return mock
			},
		},
		{
			name: helpers.FailedDelete,
			args: args{
				ctx: ctx,
				req: chatId,
			},
			err: errs.ErrDeleteChat,
			chatServiceMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := mocks.NewChatRepositoryMock(mc)
				mock.DeleteChatMock.Expect(ctx, chatId).Return(errs.ErrDeleteChat)

				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatServiceMock := tt.chatServiceMock(mc)

			service := chat.NewChatService(chatServiceMock, nil, nil, nil)

			err := service.Delete(ctx, chatId)
			require.Equal(t, tt.err, err)
		})
	}
}
