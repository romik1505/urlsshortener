package apishortener

import (
	"context"
	"github.com/jmoiron/sqlx"
	"log"
	"math/rand"
	"urlsshortener/pkg/api"
)

type GRPCServer struct {
	api.UnimplementedShortenerServer
	Db *sqlx.DB
}

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

func randomString(lenght int) string {
	str := make([]byte, lenght)
	for i := 0; i < lenght; i++ {
		str[i] = chars[rand.Intn(len(chars))]
	}
	return string(str)
}

func (g *GRPCServer) Create(ctx context.Context, in *api.Message) (*api.Message, error) {
	originalUrl := in.GetUrl()
	var shortenerUrl string
	err := g.Db.QueryRow("SELECT shortener FROM urls WHERE original=$1", originalUrl).Scan(&shortenerUrl)

	if err != nil { // Ссылка не найдена
		for i := 0; i < 3; i++ { // Решение случайной коллизии
			shortenerUrl = randomString(10)
			result, err := g.Db.Exec("INSERT INTO urls (original, shortener) VALUES ($1, $2)", originalUrl, shortenerUrl)
			if err != nil { // Рандомная строка уже присутствует в бд
				log.Println("Random generated url collision url=" + shortenerUrl)
				continue
			}
			if result != nil { // Занесено успешно
				log.Println("Successful addition: " + originalUrl + "->" + shortenerUrl)
				return &api.Message{Url: shortenerUrl}, nil
			}
		}
		return &api.Message{Url: ""}, err
	}
	log.Println("Already exist: " + originalUrl + "->" + shortenerUrl)
	return &api.Message{Url: shortenerUrl}, nil
}

func (g *GRPCServer) Get(ctx context.Context, in *api.Message) (*api.Message, error) {
	shortenerUrl := in.GetUrl()
	var originalUrl string
	err := g.Db.QueryRow("SELECT original FROM urls WHERE shortener=$1", shortenerUrl).Scan(&originalUrl)
	if err != nil {
		log.Println("Not found: " + shortenerUrl)
		return &api.Message{Url: ""}, nil
	}
	log.Println("Get success: " + shortenerUrl + "->" + originalUrl)
	return &api.Message{Url: originalUrl}, nil
}
