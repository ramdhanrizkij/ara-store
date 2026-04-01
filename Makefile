APP_NAME=ara-store

run:
	go run cmd/server/main.go

build:
	go build -o bin/app cmd/server/main.go

worker:
	go run cmd/worker/main.go

test:
	go test ./...

migrate-up:
	goose -dir migrations postgres "postgres://postgres:secret@localhost:5432/ara_store?sslmode=disable" up

migrate-down:
	goose -dir migrations postgres "postgres://postgres:secret@localhost:5432/ara_store?sslmode=disable" down

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down