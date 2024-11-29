package main

import (
	"chat-server/internal/config"
	"chat-server/internal/config/env"
	"chat-server/internal/database"
	"chat-server/internal/repository/chat"
	"chat-server/internal/service"
	deps "chat-server/pkg/chat_v1"
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", ".env", "path to config file")
	flag.Parse()
}

func main() {
	err := config.Load(configPath)
	if err != nil {
		log.Println(err.Error())
	}

	pgConfig, err := env.NewPgConfigSearcher().Get()
	if err != nil {
		log.Fatal(err.Error())
	}

	loggerConfig, err := env.NewLogConfigSearcher().Get()
	if err != nil {
		log.Fatal(err.Error())
	}

	logger := loggerConfig.SetUp()

	grpcConfig, err := env.NewGRPCConfigSearcher().Get()
	if err != nil {
		log.Fatal(err.Error())
	}

	ctx := context.Background()

	pool, closer, err := database.InitPostgresConnection(ctx, pgConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer closer()

	repository := chat.NewChatRepository(pool, logger)
	chatService := service.NewChatService(repository, logger)

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatal(err.Error())
	}

	server := grpc.NewServer()

	reflection.Register(server)
	deps.RegisterChatServer(server, chatService)

	log.Println("server starts...")
	log.Fatal(server.Serve(lis))
}
