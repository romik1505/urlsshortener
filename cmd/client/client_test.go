package main

import (
	"context"
	"google.golang.org/grpc"
	"testing"
	"urlsshortener/cmd/config"
	"urlsshortener/pkg/api"
)

func TestGRPC(t *testing.T) {
	conf := config.GetConfig()

	conn, err := grpc.Dial(conf.Address, grpc.WithInsecure())
	if err != nil {
		t.Errorf("Bad connection: %s\n", err)
	}
	cl := api.NewShortenerClient(conn)

	originalsUrl := []string{"www.yandex.ru", "www.google.com", "www.ozon.ru", "www.github.com"}
	shortenerUrl := []string{}

	for _, url := range originalsUrl {
		mess, err := cl.Create(context.Background(), &api.Message{Url: url})
		if err != nil {
			t.Fatal(err)
		}
		shortenerUrl = append(shortenerUrl, mess.GetUrl())
	}

	for i, shUrl := range shortenerUrl {
		mess, err := cl.Get(context.Background(), &api.Message{Url: shUrl})
		if err != nil {
			t.Error(err)
		}
		if mess.GetUrl() != originalsUrl[i] {
			t.Error("Error bad url relation")
		}
	}

}
