package chat

import (
	"chat-server/internal/model"
	"chat-server/pkg/errs"
	"context"
)

func (s *ChatServiceImpl) Create(ctx context.Context, usernames []model.User) (int64, error) {
	var chatId int64

	err := s.txManeger.ReadCommitted(ctx, func(ctx context.Context) error {
		var txErr error

		chatId, txErr = s.chatRepository.CreateChat(ctx)
		if txErr != nil {
			return errs.ErrCreateChat
		}

		txErr = s.memberRepository.CreateMember(ctx, &model.CreateMemberRequest{
			ChatId:    chatId,
			Usernames: usernames,
		})
		if txErr != nil {
			return errs.ErrCreateMember
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return chatId, nil
}
