build:
	docker-compose build urlshortener

run:
	docker-compose up urlshortener

test:
	go test -v ./cmd/client/client_test.go

server:
	go run ./cmd/server/main.go

migrate:
	migrate -path ./schema -database 'postgres://postgres:12345678@localhost:5432/urlshortener?sslmode=disable' up