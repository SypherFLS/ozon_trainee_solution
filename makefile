.PHONY: run build test docker-up docker-up-d docker-down docker-logs docker-build 

run:
	cd backend && go run ./cmd/main.go

build:
	cd backend && mkdir -p ./bin && go build -o ./bin/main ./cmd/main.go

test:
	cd backend && go test -count=1 ./... -v

docker-up:
	docker-compose up

docker-up-d:
	docker-compose up -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f

docker-build:
	docker-compose build

