package app

import (
	"chat-server/internal/api/chat"
	"chat-server/internal/client/database"
	"chat-server/internal/client/database/postgres"
	"chat-server/internal/client/database/transaction"
	"chat-server/internal/closer"
	"chat-server/internal/config"
	"chat-server/internal/config/env"
	"chat-server/internal/repository"
	chatRepository "chat-server/internal/repository/chat"
	"chat-server/internal/repository/chat_members"
	"chat-server/internal/repository/messages"
	"chat-server/internal/service"
	chatService "chat-server/internal/service/chat"
	"context"
	"log"
)

type serviceProvider struct {
	pgConfig   *config.PgConfig
	grpcConfig *config.GRPCConfig

	dbClient  database.Client
	txManeger database.TxManeger

	chatRepository    repository.ChatRepository
	messageRepository repository.MessageRepository
	membersRepository repository.MembersRepository

	chatService service.ChatService
	chatImpl    *chat.ChatApiImpl
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() *config.PgConfig {
	if s.pgConfig == nil {
		cfgSearcher := env.NewPgCfgSearcher()

		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() *config.GRPCConfig {
	if s.grpcConfig == nil {
		cfgSearcher := env.NewGRPCCfgSearcher()

		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) database.Client {
	if s.dbClient == nil {
		dbc, err := postgres.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create DBClient: %s", err.Error())
		}

		err = dbc.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("failed to ping database: %s", err.Error())
		}

		closer.Add(dbc.Close)

		s.dbClient = dbc
	}
	return s.dbClient
}

func (s *serviceProvider) TxManeger(ctx context.Context) database.TxManeger {
	if s.txManeger == nil {
		s.txManeger = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManeger
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepository.NewChatRepositoryImpl(s.DBClient(ctx))
	}

	return s.chatRepository
}

func (s *serviceProvider) MessageRepository(ctx context.Context) repository.MessageRepository {
	if s.messageRepository == nil {
		s.messageRepository = messages.NewMessagesRepository(s.DBClient(ctx))
	}

	return s.messageRepository
}

func (s *serviceProvider) MemberRepository(ctx context.Context) repository.MembersRepository {
	if s.membersRepository == nil {
		s.membersRepository = chat_members.NewMembersRepository(s.DBClient(ctx))
	}

	return s.membersRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewChatService(
			s.ChatRepository(ctx),
			s.MemberRepository(ctx),
			s.MessageRepository(ctx),
			s.TxManeger(ctx))
	}

	return s.chatService
}

func (s *serviceProvider) ChatImpl(ctx context.Context) *chat.ChatApiImpl {
	if s.chatImpl == nil {
		s.chatImpl = chat.NewChatApiImpl(s.ChatService(ctx))
	}

	return s.chatImpl
}
