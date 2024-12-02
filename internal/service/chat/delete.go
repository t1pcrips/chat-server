package chat

import (
	"chat-server/pkg/errs"
	"context"
)

func (s *ChatServiceImpl) Delete(ctx context.Context, chatId int64) error {
	err := s.chatRepository.DeleteChat(ctx, chatId)
	if err != nil {
		return errs.ErrDeleteChat
	}

	return nil
}
