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
	go run cmd/migrate/main.go up

migrate-down:
	go run cmd/migrate/main.go down
	
docker-up:
	docker-compose up -d

docker-down:
	docker-compose down