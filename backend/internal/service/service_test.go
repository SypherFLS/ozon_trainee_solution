package service

import (
	"ots/internal/shortener"
	"ots/internal/storage/memory"
	"testing"
)

var Data = []string{"https://google.com", "https://google.ru", "https://google.dot", "https://google.com"}

func TestPipeline(t *testing.T) {
	repo := memory.New()
	gen := shortener.NewGen()
	svc := New(repo, gen)

	for _, url := range Data {
		short := shortenHelper(t, svc, url)
		t.Log(url, "->", short)
	}
}

func shortenHelper(t *testing.T, svc *Service, url string) string {
	t.Helper()

	short, err := svc.Shorten(url)
	if err != nil {
		t.Fatal(err)
	}

	return short
}