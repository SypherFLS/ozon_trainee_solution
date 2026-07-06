.PHONY: run build test

run:
	cd backend && go run ./cmd/main.go

build:
	cd backend && mkdir -p ./bin && go build -o ./bin/main ./cmd/main.go

test:
	cd backend && go test -count=1 ./... -v
