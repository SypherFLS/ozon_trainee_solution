package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"ots/internal/api/handler"
	"ots/internal/api/router"
	"ots/internal/config"
	"ots/internal/repository"
	"ots/internal/service"
	"ots/internal/shortener"
	"ots/internal/storage/memory"
	"ots/internal/storage/postgres"
)

func main() {
	cfg := config.MustLoad()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	var repo repository.Repository

	switch cfg.Storage {
	case "memory":
		repo = memory.New()
	case "postgres":
		db, err := postgres.InitDB(cfg)
		if err != nil {
			log.Fatal(err)
		}
		repo = postgres.New(db)
	default:
		log.Fatal("unknown storage: ", cfg.Storage)
	}

	gen := shortener.NewGen()
	svc := service.New(repo, gen)
	h := handler.New(svc, logger)
	r := router.New(h)

	addr := fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port)

	logger.Info("starting server", "addr", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}