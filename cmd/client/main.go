package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/fatih/color"
	desc "github.com/t1pcrips/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	addres = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(addres, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	if err != nil {
		log.Fatal(fmt.Sprintf(color.RedString("failed to connected to server - %d, error - %v", addres, err)))
	}

	c := desc.NewChatClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := c.Create(ctx, &desc.CreateRequest{
		Usernames: []string{
			gofakeit.Name(), gofakeit.Name(), gofakeit.Name(), gofakeit.Name(),
		},
	})
	if err != nil {
		log.Fatalf(fmt.Sprintf(color.RedString("failed to create Request error - %v\n", err)))
	}

	log.Printf(color.RedString("Note Info:\n") + color.GreenString("%+v", r.GetId()))
}
