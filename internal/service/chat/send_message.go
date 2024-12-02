package chat

import (
	"chat-server/internal/model"
	"chat-server/pkg/errs"
	"context"
	"fmt"
)

func (s *ChatServiceImpl) SendMessage(ctx context.Context, info *model.CreateMessageRequest) error {
	err := s.txManeger.ReadCommitted(ctx, func(ctx context.Context) error {
		existsChat, txErrChat := s.chatRepository.CheckChatExists(ctx, info.ChatId)
		if txErrChat != nil {
			return txErrChat
		}
		if !existsChat {
			return errs.ErrChatNotFound
		}

		existsMember, txErrMember := s.memberRepository.CheckMemberExists(ctx, info.From)
		if txErrMember != nil {
			return txErrMember
		}
		if !existsMember {
			return errs.ErrMemberNotFound
		}

		txErr := s.messageRepository.CreateMessage(ctx, info)
		if txErr != nil {
			return errs.ErrCreateMember
		}
		return nil
	})
	if err != nil {
		return err
	}
	fmt.Println("yes")
	return nil
}
