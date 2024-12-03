package tests

import (
	"chat-server/internal/api/chat"
	"chat-server/internal/model"
	"chat-server/internal/service"
	"chat-server/internal/service/mocks"
	desc "chat-server/pkg/chat_v1"
	"chat-server/pkg/errs"
	"context"
	"fmt"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand/v2"
	"testing"
)

func TestCreate(t *testing.T) {
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		chatId    = rand.Int64()
		usernames = names(rand.IntN(5))

		req = &desc.CreateRequest{
			Usernames: usernames,
		}

		info = &model.CreateChatRequest{
			Usernames: usernames,
		}

		resp = &desc.CreateResponse{
			Id: chatId,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: resp,
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := mocks.NewChatServiceMock(mc)
				mock.CreateMock.Expect(ctx, info).Return(chatId, nil)
				return mock
			},
		},
		{
			name: "service error mock",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  status.Error(codes.NotFound, errs.ErrFailedService.Error()),
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := mocks.NewChatServiceMock(mc)
				mock.CreateMock.Expect(ctx, info).Return(0, errs.ErrFailedService)

				return mock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatServiceMock := tt.chatServiceMock(mc)
			api := chat.NewChatApiImpl(chatServiceMock)

			testResponse, err := api.Create(tt.args.ctx, tt.args.req)
			if err != nil {
				fmt.Println(tt.err, err)
			}
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, testResponse)
		})
	}

}
