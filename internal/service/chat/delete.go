package chat

import (
	"context"
	"github.com/t1pcrips/chat-service/pkg/errs"
)

func (s *ChatServiceImpl) Delete(ctx context.Context, chatId int64) error {
	err := s.chatRepository.DeleteChat(ctx, chatId)
	if err != nil {
		return errs.ErrDeleteChat
	}

	return nil
}
