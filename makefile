.PHONY: run build test test-cover test-short test-bench docker-up docker-up-d docker-down docker-logs docker-build 

run:
	cd backend && go run ./cmd/main.go

build:
	cd backend && mkdir -p ./bin && go build -o ./bin/main ./cmd/main.go

test:
	cd backend && go test -count=1 ./internal/... -v

test-cover:
	cd backend && go test -count=1 ./internal/... -cover

test-short:
	cd backend && go test -count=1 -short ./internal/... -v

test-bench:
	cd backend && go test -bench=. -benchmem ./internal/...

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

