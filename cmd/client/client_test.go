package main

import (
	"UrlShortener/cmd/config"
	"UrlShortener/pkg/api"
	"context"
	"google.golang.org/grpc"
	"testing"
)

func TestGRPC(t *testing.T)  {
	conf, err := config.LoadConfig("../../config/")
	if err != nil {
		t.Errorf("Error load configuration: %s\n", err)
	}

	conn, err := grpc.Dial(conf.Server.Address, grpc.WithInsecure())
	if err != nil {
		t.Errorf("Bad connection: %s\n", err)
	}
	cl := api.NewShortenerClient(conn)

	originalsUrl := []string{"ABOBA", "AMOGUS", "ABIBUS", "NEW"}
	shortenerUrl := []string{}

	for _, url := range originalsUrl{
		mess, err := cl.Create(context.Background(), &api.Message{Url: url})
		if err != nil {
			t.Error(err)
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