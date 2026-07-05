package postgres

import (
	"errors"

	"ots/internal/apperrors"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) Save(original, short string) error {
	var existing Link
	err := s.db.Where("original = ?", original).First(&existing).Error
	if err == nil {
		return nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	link := Link{Original: original, Short: short}
	if err := s.db.Create(&link).Error; err != nil {
		if !isDuplicateKey(err) {
			return err
		}

		var byOriginal Link
		if s.db.Where("original = ?", original).First(&byOriginal).Error == nil {
			return nil
		}

		return apperrors.ErrConflict
	}

	return nil
}

func (s *Storage) GetByOriginal(original string) (string, error) {
	var link Link
	err := s.db.Where("original = ?", original).First(&link).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", apperrors.ErrNotFound
	}
	if err != nil {
		return "", err
	}

	return link.Short, nil
}

func (s *Storage) GetByShort(short string) (string, error) {
	var link Link
	err := s.db.Where("short = ?", short).First(&link).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", apperrors.ErrNotFound
	}
	if err != nil {
		return "", err
	}

	return link.Original, nil
}

func (s *Storage) ExistsShort(short string) bool {
	var count int64
	s.db.Model(&Link{}).Where("short = ?", short).Limit(1).Count(&count)
	return count > 0
}

func isDuplicateKey(err error) bool {
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return true
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return true
	}

	return false
}
