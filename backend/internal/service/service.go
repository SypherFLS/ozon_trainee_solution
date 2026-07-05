package service

import (
	"errors"

	"ots/internal/apperrors"
	"ots/internal/repository"
	"ots/internal/shortener"
)

type Service struct {
    repo repository.Repository
    gen  *shortener.Generator
}

func New(repo repository.Repository, gen *shortener.Generator) *Service {
	return &Service{
		repo: repo,
		gen:  gen,
	}
}

func (s *Service) GetOriginal(short string) (string, error) {
    return s.repo.GetByShort(short)
}

func (s *Service) Shorten(url string) (string, error) {
	short, err := s.repo.GetByOriginal(url)
	if err == nil {
		return short, nil
	}
	if !errors.Is(err, apperrors.ErrNotFound) {
		return "", err
	}

	for {
		short, err := s.gen.Generate()
		if err != nil {
			return "", err
		}
		if !s.repo.ExistsShort(short) {
			if err := s.repo.Save(url, short); err != nil {
				if errors.Is(err, apperrors.ErrConflict) {
					continue
				}
				return "", err
			}
			short, err = s.repo.GetByOriginal(url)
			if err == nil {
				return short, nil
			}
			if !errors.Is(err, apperrors.ErrNotFound) {
				return "", err
			}
		}
	}
}	