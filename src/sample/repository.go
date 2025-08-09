package sample

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

type SampleRepo interface {
	CreateSample() error
}

func (s *Repository) CreateSample() error {
	// actual implementation
	return nil
}