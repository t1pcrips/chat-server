package tests

import (
	"context"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"github.com/t1pcrips/chat-service/internal/model"
	"github.com/t1pcrips/chat-service/internal/repository"
	mocksRepo "github.com/t1pcrips/chat-service/internal/repository/mocks"
	"github.com/t1pcrips/chat-service/internal/service/chat"
	"github.com/t1pcrips/chat-service/pkg/helpers"
	"github.com/t1pcrips/platform-pkg/pkg/database"
	"github.com/t1pcrips/platform-pkg/pkg/database/transaction"
	"math/rand"
	"testing"
)

type an interface{}

func TestCreate(t *testing.T) {
	type chatRepositoryMocksFunc func(mc *minimock.Controller) repository.ChatRepository
	type membersRepositoryMocksFunc func(mc *minimock.Controller) repository.MembersRepository
	type txManegerMocksFunc func(mc *minimock.Controller) database.TxManeger

	type args struct {
		ctx context.Context
		req *model.CreateChatRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		chatId    = rand.Int63()
		usernames = helpers.Names(rand.Intn(10))

		chatReq = &model.CreateChatRequest{
			Usernames: usernames,
		}

		memberReq = &model.CreateMemberRequest{
			ChatId:    chatId,
			Usernames: usernames,
		}
		txFaker, ctxWithTx = helpers.TxFakerAndCtxWithSetup(ctx, t, true)
		transactorMock     = helpers.FakeTransactorMock(ctxWithTx, t, txFaker)
	)

	tests := []struct {
		name                   string
		args                   args
		want                   int64
		err                    error
		chatRepositoryMocks    chatRepositoryMocksFunc
		membersRepositoryMocks membersRepositoryMocksFunc
		txManegerMocks         txManegerMocksFunc
	}{
		{
			name: helpers.SuccessCreate,
			args: args{
				ctx: ctxWithTx,
				req: chatReq,
			},
			want: chatId,
			err:  nil,
			chatRepositoryMocks: func(mc *minimock.Controller) repository.ChatRepository {
				mock := mocksRepo.NewChatRepositoryMock(mc)
				mock.CreateChatMock.Expect(ctxWithTx).Return(chatId, nil)
				return mock
			},
			membersRepositoryMocks: func(mc *minimock.Controller) repository.MembersRepository {
				mock := mocksRepo.NewMembersRepositoryMock(mc)
				mock.CreateMemberMock.Expect(ctxWithTx, memberReq).Return(nil)
				return mock
			},
			txManegerMocks: func(mc *minimock.Controller) database.TxManeger {
				txManegerMock := transaction.NewTransactionManager(transactorMock)
				return txManegerMock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatRepositoryMocks := tt.chatRepositoryMocks(mc)
			memberRepositoryMocks := tt.membersRepositoryMocks(mc)
			txManegerMocks := tt.txManegerMocks(mc)
			service := chat.NewChatService(chatRepositoryMocks, memberRepositoryMocks, nil, txManegerMocks)

			response, err := service.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.want, response)
			require.Equal(t, tt.err, err)
		})
	}
}
