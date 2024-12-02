package main

import (
	"chat-server/internal/app"
	"context"
	"flag"
	"log"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", ".env", "path to config file")
	flag.Parse()
}

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx, configPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}
