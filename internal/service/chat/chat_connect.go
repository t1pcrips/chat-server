package chat

import (
	"context"
	"errors"
	"github.com/t1pcrips/chat-service/internal/model"
	"github.com/t1pcrips/chat-service/internal/service"
	"github.com/t1pcrips/chat-service/internal/service/chat/streams"
	"github.com/t1pcrips/chat-service/pkg/errs"
)

func (s *ChatServiceImpl) Connect(info *model.User, stream service.StreamChatMessages) error {
	err := s.txManeger.ReadCommitted(stream.Context(), func(ctx context.Context) error {
		existsChat, txErrChat := s.chatRepository.CheckChatExists(ctx, info.ChatId)
		if txErrChat != nil {
			return txErrChat
		}
		if !existsChat {
			return errs.ErrChatNotFound
		}

		existsMember, _ := s.memberRepository.CheckMemberExists(ctx, info.Username)
		//if txErrMember != nil {
		//	return txErrMember
		//}

		if !existsMember && existsChat {
			err := s.memberRepository.CreateMember(ctx, &model.CreateMemberRequest{
				ChatId:    info.ChatId,
				Usernames: []string{info.Username},
			})

			if err != nil {
				return err
			}

			return nil
		}

		return nil
	})
	if err != nil {
		if !errors.Is(err, errs.ErrExec) {
			return err
		}
	}

	// получаем канал для чата если его нет то инициализируем через эти каналы будут пердадваться сообщения польщователей чата
	chanels, ok := s.chatsMessageChannels.GetChanelForChat(info.ChatId)
	if !ok {
		chanels = s.chatsMessageChannels.InitChannelForChat(info.ChatId, 100)
	}

	// получаем чат для чат id
	chat := s.chats.CreateOrGetChat(info.ChatId)
	chat.SetStreamForUser(info, stream)

	err = s.chatProcessingMessages(info, chat, chanels, stream)
	if err != nil {
		return err
	}

	return nil
}

func (s *ChatServiceImpl) chatProcessingMessages(
	user *model.User,
	chat *streams.Chat,
	chanChat chan *model.Message,
	stream service.StreamChatMessages,
) error {

	for {
		select {
		case msg, okMsg := <-chanChat:
			if !okMsg {
				return nil
			}

			for _, st := range chat.GetStreamForUsers() {
				err := st.Send(msg)
				if err != nil {
					return err
				}
			}
		case <-stream.Context().Done():
			chat.DeleteStreamForUser(user)
			return nil
		}
	}
}
