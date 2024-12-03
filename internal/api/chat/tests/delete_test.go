package tests

import (
	"chat-server/internal/api/chat"
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
	"math/rand"
	"testing"
)

func TestDelete(t *testing.T) {
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *desc.DeleteRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		chatId = rand.Int63()

		req = &desc.DeleteRequest{
			Id: chatId,
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
			name: successDelete,
			args: args{
				ctx: ctx,
				req: req,
			},
			want: &emptypb.Empty{},
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := mocks.NewChatServiceMock(mc)
				mock.DeleteMock.Expect(ctx, chatId).Return(nil)
				return mock
			},
		},
		{
			name: failedDelete,
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  status.Error(codes.NotFound, errs.ErrFailedService.Error()),
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := mocks.NewChatServiceMock(mc)
				mock.DeleteMock.Expect(ctx, chatId).Return(errs.ErrFailedService)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatServiceMock := tt.chatServiceMock(mc)
			api := chat.NewChatApiImpl(chatServiceMock)

			testResponse, err := api.Delete(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.want, testResponse)
			require.Equal(t, tt.err, err)
		})
	}
}
