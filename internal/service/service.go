package service

import (
	"chat-server/internal/repository"
	"chat-server/internal/repository/chat"
	deps "chat-server/pkg/chat_v1"
	"context"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ChatService struct {
	deps.UnimplementedChatServer
	logger     *zerolog.Logger
	repository *chat.ChatRepository
}

func NewChatService(repository *chat.ChatRepository, logger *zerolog.Logger) *ChatService {
	return &ChatService{
		logger:     logger,
		repository: repository,
	}
}

func (s *ChatService) Create(ctx context.Context, req *deps.CreateRequest) (*deps.CreateResponse, error) {
	s.logRequest("CREATE CHAT", req)

	chatId, err := s.repository.CreateChat(ctx, req.GetUsernames())
	if err != nil {
		s.logger.Err(err).Msg("failed to create chat")
		return nil, status.Error(codes.NotFound, "failed to create chat")
	}

	return &deps.CreateResponse{
		Id: chatId,
	}, nil
}

func (s *ChatService) Delete(ctx context.Context, req *deps.DeleteRequest) (*emptypb.Empty, error) {
	s.logRequest("DELETE CHAT", req)

	err := s.repository.DeleteChat(ctx, req.GetId())
	if err != nil {
		s.logger.Err(err).Msg("failed to delete chat")
		return nil, status.Error(codes.NotFound, "failed to delete user with this id")
	}

	return &emptypb.Empty{}, nil
}

func (s *ChatService) SendMessage(ctx context.Context, req *deps.SendMessageRequest) (*emptypb.Empty, error) {
	s.logRequest("SEND MESSAGE", req)

	err := s.repository.SendMessage(ctx, &repository.SendMessageRequest{
		From:     req.GetFrom(),
		Text:     req.GetText(),
		TimeSend: req.GetTimestamp().AsTime(),
	})

	if err != nil {
		s.logger.Err(err).Msg("failed to send message")
		return nil, status.Error(codes.NotFound, "failed to delete")
	}

	return &emptypb.Empty{}, nil
}

func (s *ChatService) logRequest(method string, req interface{}) {
	s.logger.Debug().
		Str("method", method).
		Interface("request", req).Msg("try to process user request")
}
