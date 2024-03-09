package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	desc "github.com/t1pcrips/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	grpcPort = 50051
)

type server struct {
	desc.UnimplementedChatServer
}

// Create . . .
func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf(fmt.Sprintf(color.HiCyanString("Create Message User: %v", req.Usernames)))
	return &desc.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

// Deleate . . .
func (s *server) Deleate(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf(fmt.Sprintf(color.HiCyanString("Deleate Message ID:", req.Id)))
	return &emptypb.Empty{}, nil
}

// SendMessage . . .
func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf(fmt.Sprintf(color.HiCyanString("Send Message at - %v", req.Timestamp)))
	return &emptypb.Empty{}, nil
}

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatal(fmt.Sprintf(color.RedString("failed to listen on port - %d, error - %v", grpcPort, err)))
	}

	s := grpc.NewServer()
	reflection.RegisterV1(s)
	desc.RegisterChatServer(s, &server{})

	log.Print(fmt.Sprintf(color.HiMagentaString("server listening at - %v", listen.Addr())))

	if err := s.Serve(listen); err != nil {
		log.Fatalf(fmt.Sprintf(color.RedString("failed to serve - %v", err)))

	}
}
