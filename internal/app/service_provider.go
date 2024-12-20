package app

import (
	"context"
	"github.com/t1pcrips/chat-service/internal/api/chat"
	"github.com/t1pcrips/chat-service/internal/client"
	"github.com/t1pcrips/chat-service/internal/client/access"
	"github.com/t1pcrips/chat-service/internal/config"
	"github.com/t1pcrips/chat-service/internal/config/env"
	"github.com/t1pcrips/chat-service/internal/interceptor"
	"github.com/t1pcrips/chat-service/internal/repository"
	chatRepository "github.com/t1pcrips/chat-service/internal/repository/chat"
	"github.com/t1pcrips/chat-service/internal/repository/chat_members"
	"github.com/t1pcrips/chat-service/internal/repository/messages"
	"github.com/t1pcrips/chat-service/internal/service"
	chatService "github.com/t1pcrips/chat-service/internal/service/chat"
	"github.com/t1pcrips/chat-service/internal/service/chat/streams"
	"github.com/t1pcrips/chat-service/pkg/access_v1"
	"github.com/t1pcrips/platform-pkg/pkg/closer"
	"github.com/t1pcrips/platform-pkg/pkg/database"
	"github.com/t1pcrips/platform-pkg/pkg/database/postgres"
	"github.com/t1pcrips/platform-pkg/pkg/database/transaction"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type serviceProvider struct {
	pgConfig          *config.PgConfig
	grpcConfig        *config.GRPCConfig
	httpConfig        *config.HTTPConfig
	authConfig        *config.AuthConfig
	swaggerConfig     *config.SWAGGERConfig
	accessInterceptor *interceptor.AccessInterceptor

	dbClient  database.Client
	txManeger database.TxManeger

	chatRepository    repository.ChatRepository
	messageRepository repository.MessageRepository
	membersRepository repository.MembersRepository

	accessInterseptor interceptor.AccessInterceptor

	chats                *streams.Chats
	chatsMessageChannels *streams.ChatMessageChannels

	accessClient   client.AccessClient
	accessV1Client access_v1.AccessClient
	chatService    service.ChatService
	chatImpl       *chat.ChatApiImpl
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() *config.PgConfig {
	if s.pgConfig == nil {
		cfgSearcher := env.NewPgConfigSearcher()

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
		cfgSearcher := env.NewGRPCConfigSearcher()

		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) HTTPConfig() *config.HTTPConfig {
	if s.httpConfig == nil {
		cfgSearcher := env.NewHTTPConfigSearcher()

		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) SWAGGERConfig() *config.SWAGGERConfig {
	if s.swaggerConfig == nil {
		cfgSearcher := env.NewSwaggerConfigSearcher()

		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("failed to get swagger config: %s", err.Error())
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
}

func (s *serviceProvider) AUTHConfig() *config.AuthConfig {
	if s.authConfig == nil {
		cfgSearcher := env.NewAuthConfigSearcher()

		cfg, err := cfgSearcher.Get()
		if err != nil {
			return nil
		}

		s.authConfig = cfg
	}

	return s.authConfig
}

func (s *serviceProvider) AccessInterceptor(ctx context.Context) *interceptor.AccessInterceptor {
	if s.accessInterceptor == nil {
		s.accessInterceptor = interceptor.NewAccessInterceptor(s.AccessClient(ctx))
	}

	return s.accessInterceptor
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

func (s *serviceProvider) AccessV1Client(ctx context.Context) access_v1.AccessClient {
	if s.accessV1Client == nil {
		conn, err := grpc.NewClient(
			s.AUTHConfig().Address(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Fatalf("failed to connect to auth service: %s", err.Error())
		}

		s.accessV1Client = access_v1.NewAccessClient(conn)
	}

	return s.accessV1Client
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
			s.TxManeger(ctx),
			s.Chats(),
			s.ChatsMessageChannels(),
		)
	}

	return s.chatService
}

func (s *serviceProvider) ChatImpl(ctx context.Context) *chat.ChatApiImpl {
	if s.chatImpl == nil {
		s.chatImpl = chat.NewChatApiImpl(s.ChatService(ctx))
	}

	return s.chatImpl
}

func (s *serviceProvider) AccessClient(ctx context.Context) client.AccessClient {
	if s.accessClient == nil {

		s.accessClient = access.NewAccessClientImpl(s.AccessV1Client(ctx))
	}
	return s.accessClient
}

func (s *serviceProvider) ChatsMessageChannels() *streams.ChatMessageChannels {
	if s.chatsMessageChannels == nil {
		s.chatsMessageChannels = streams.NewChatMessageChannels()
	}

	return s.chatsMessageChannels
}

func (s *serviceProvider) Chats() *streams.Chats {
	if s.chats == nil {
		s.chats = streams.NewChats()
	}

	return s.chats
}
