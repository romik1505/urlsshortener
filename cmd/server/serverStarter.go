package main

import (
	"UrlShortener/cmd/config"
	"UrlShortener/pkg/api"
	us "UrlShortener/pkg/shortener"
	"database/sql"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"time"
)

func init()  {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	conf, err := config.LoadConfig("config/")
	if err != nil {
		log.Fatal(err)
	}

	connStr := conf.GetDbConnectionString()
	db, err := sql.Open(conf.Db.Drivername, connStr)
	defer db.Close()

	if err != nil {
		log.Panic(err)
	}

	s := grpc.NewServer()
	srv := &us.GRPCServer{Db: db}
	api.RegisterShortenerServer(s, srv)

	lis, err := net.Listen(conf.Server.Network, conf.Server.Address)
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
