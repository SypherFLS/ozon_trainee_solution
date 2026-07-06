package service

import (
	"errors"
	"ots/internal/apperrors"
	"ots/internal/shortener"
	"ots/internal/storage/memory"
	"testing"
)

var Data = []string{"https://google.com", "https://google.ru", "https://google.dot", "https://google.com"}

func TestPipeline(t *testing.T) {
	repo := memory.New()
	gen := shortener.NewGen()
	svc := New(repo, gen)
	result := []string{}
	for _, url := range Data {
		short := shortenHelper(t, svc, url)
		result = append(result, short)
		t.Log(url, "->", short)
	}

	if result[0] != result[3] {
		t.Logf("expected %s == %s", result[0], result[3])
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

func TestShortenNewURL(t *testing.T) {

	repo := memory.New()
	gen := shortener.NewGen()
	svc := New(repo, gen)

	url := "https://google.com"

	short, err := svc.Shorten(url)

	if err != nil {
		t.Fatalf("Shorten() error = %v", err)
	}
	if short == "" {
		t.Error("Shorten() returned empty string")
	}
}

func TestGetOriginal(t *testing.T) {
	repo := memory.New()
	gen := shortener.NewGen()
	svc := New(repo, gen)

	originalURL := "https://google.com"
	shortURL, _ := svc.Shorten(originalURL)

	retrieved, err := svc.GetOriginal(shortURL)

	if err != nil {
		t.Fatalf("GetOriginal() error = %v", err)
	}
	if retrieved != originalURL {
		t.Errorf("GetOriginal() got %s, want %s", retrieved, originalURL)
	}
}

func TestShortenExistingURL(t *testing.T) {
	repo := memory.New()
	gen := shortener.NewGen()
	svc := New(repo, gen)

	url := "https://google.com"
	short1, _ := svc.Shorten(url)

	short2, err := svc.Shorten(url)

	if err != nil {
		t.Fatalf("Shorten() error = %v", err)
	}
	if short1 != short2 {
		t.Errorf("Shorten() same URL got %s, want %s", short2, short1)
	}
}

func TestGetOriginalNotFound(t *testing.T) {
	repo := memory.New()
	gen := shortener.NewGen()
	svc := New(repo, gen)

	_, err := svc.GetOriginal("nonexistent")

	if !errors.Is(err, apperrors.ErrNotFound) {
		t.Errorf("GetOriginal() error = %v, want ErrNotFound", err)
	}
}

func TestShortenMultipleDifferentURLs(t *testing.T) {
	repo := memory.New()
	gen := shortener.NewGen()
	svc := New(repo, gen)

	urls := []string{
		"https://google.com",
		"https://yandex.ru",
		"https://github.com",
		"https://stackoverflow.com",
	}

	shorts := make(map[string]string)

	for _, url := range urls {
		short, err := svc.Shorten(url)
		if err != nil {
			t.Fatalf("Shorten(%s) error = %v", url, err)
		}
		shorts[url] = short
	}

	if len(shorts) != len(urls) {
		t.Errorf("Shorten() got %d unique URLs, want %d", len(shorts), len(urls))
	}

	for url, short := range shorts {
		retrieved, _ := svc.GetOriginal(short)
		if retrieved != url {
			t.Errorf("Shorten() URL mismatch: got %s, want %s", retrieved, url)
		}
	}
}
