package chat

import (
	"github.com/t1pcrips/chat-service/internal/model"
	"github.com/t1pcrips/chat-service/pkg/errs"

	"context"
)

func (s *ChatServiceImpl) Create(ctx context.Context, info *model.CreateChatRequest) (int64, error) {
	var chatId int64

	err := s.txManeger.ReadCommitted(ctx, func(ctx context.Context) error {
		var txErr error

		chatId, txErr = s.chatRepository.CreateChat(ctx)
		if txErr != nil {
			return errs.ErrCreateChat
		}

		txErr = s.memberRepository.CreateMember(ctx, &model.CreateMemberRequest{
			ChatId:    chatId,
			Usernames: info.Usernames,
		})
		if txErr != nil {
			return errs.ErrCreateMember
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	s.chatsMessageChannels.InitChannelForChat(chatId, 100)

	return chatId, nil
}
