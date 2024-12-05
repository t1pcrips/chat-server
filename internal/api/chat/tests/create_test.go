package tests

import (
	"context"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"github.com/t1pcrips/chat-service/internal/api/chat"
	"github.com/t1pcrips/chat-service/internal/model"
	"github.com/t1pcrips/chat-service/internal/service"
	"github.com/t1pcrips/chat-service/internal/service/mocks"
	desc "github.com/t1pcrips/chat-service/pkg/chat_v1"
	"github.com/t1pcrips/chat-service/pkg/errs"
	"github.com/t1pcrips/chat-service/pkg/helpers"
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
		usernames = helpers.Names(rand.IntN(5))

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
			name: helpers.SuccessCreate,
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
			name: helpers.FailedCreate,
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
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, testResponse)
		})
	}

}
