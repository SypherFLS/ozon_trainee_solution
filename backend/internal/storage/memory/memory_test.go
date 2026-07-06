package memory

import (
	"errors"
	"ots/internal/apperrors"
	"testing"
)

func TestSaveAndGet(t *testing.T) {
	storage := New()
	original := "https://google.com"
	short := "abc123"

	err := storage.Save(original, short)
	if err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	result, err := storage.GetByOriginal(original)
	if err != nil {
		t.Fatalf("GetByOriginal() error = %v", err)
	}
	if result != short {
		t.Errorf("GetByOriginal() got %s, want %s", result, short)
	}
}

func TestSaveDuplicate(t *testing.T) {
	storage := New()
	original := "https://google.com"
	short := "abc123"

	storage.Save(original, short)
	err := storage.Save(original, short)

	if err != nil {
		t.Errorf("Save() duplicate should return nil, got %v", err)
	}
	t.Logf("duplicate save did not cause an error, as expected %v", err)
}

func TestGetNotFound(t *testing.T) {
	storage := New()

	_, err := storage.GetByOriginal("nonexistent")

	if !errors.Is(err, apperrors.ErrNotFound) {
		t.Errorf("GetByOriginal() got %v, want ErrNotFound", err)
	}
}

func TestSaveConflictShort(t *testing.T) {
	storage := New()
	original1 := "https://google.com"
	original2 := "https://yandex.ru"
	short := "abc123"

	storage.Save(original1, short)
	err := storage.Save(original2, short)

	if !errors.Is(err, apperrors.ErrConflict) {
		t.Errorf("Save() conflict got error = %v, want ErrConflict", err)
	}
}

func TestGetByShort_Found(t *testing.T) {
	storage := New()
	original := "https://google.com"
	short := "abc123"
	storage.Save(original, short)

	result, err := storage.GetByShort(short)

	if err != nil {
		t.Fatalf("GetByShort() error = %v, want nil", err)
	}
	if result != original {
		t.Errorf("GetByShort() got %s, want %s", result, original)
	}
}

func TestGetByShort_NotFound(t *testing.T) {
	storage := New()

	_, err := storage.GetByShort("nonexistent")

	if !errors.Is(err, apperrors.ErrNotFound) {
		t.Errorf("GetByShort() got error = %v, want ErrNotFound", err)
	}
}

func TestExistsShort_True(t *testing.T) {
	storage := New()
	original := "https://google.com"
	short := "abc123"
	storage.Save(original, short)

	exists := storage.ExistsShort(short)

	if !exists {
		t.Errorf("ExistsShort() got false, want true")
	}
}

func TestExistsShort_False(t *testing.T) {
	storage := New()

	exists := storage.ExistsShort("nonexistent")

	if exists {
		t.Errorf("ExistsShort() got true, want false")
	}
}
