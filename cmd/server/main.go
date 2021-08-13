package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"time"
	"urlsshortener/cmd/config"
	"urlsshortener/pkg/api"
	us "urlsshortener/pkg/shortener"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	conf := config.GetConfig()

	connStr := conf.GetDbConnectionString()

	db, err := sqlx.Open(conf.DbDrivername, connStr)
	defer db.Close()

	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln("Bad connection " + connStr)
	}

	s := grpc.NewServer()
	srv := &us.GRPCServer{Db: db}
	api.RegisterShortenerServer(s, srv)

	lis, err := net.Listen(conf.Network, conf.Address)
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
