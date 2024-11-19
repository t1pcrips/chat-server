package main

import (
	deps "chat-server/pkg/chat_v1"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"math/rand"
	"net"
)

type server struct {
	deps.UnimplementedChatServer
}

func (s *server) Create(ctx context.Context, req *deps.CreateRequest) (*deps.CreateResponse, error) {
	log.Println("create: ", req.Usernames)
	return &deps.CreateResponse{
		Id: rand.Int63() % 100,
	}, nil
}

func (s *server) Delete(ctx context.Context, req *deps.DeleteRequest) (*emptypb.Empty, error) {
	log.Println("delete: ", req.Id)
	return &emptypb.Empty{}, nil
}

func (s *server) SendMessage(ctx context.Context, req *deps.SendMessageRequest) (*emptypb.Empty, error) {
	log.Println("send: ", req.Text)
	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err.Error())
	}

	s := grpc.NewServer()
	reflection.Register(s)
	deps.RegisterChatServer(s, &server{})

	log.Fatal(s.Serve(lis))
}
