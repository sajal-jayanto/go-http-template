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
	Create()  error
	FindAll() error
	FindOne() error
	Update()  error
	Remove()  error
}

func (s *Repository) Create() error {
	// actual implementation
	return nil
}

func (s *Repository) FindAll() error {
	// actual implementation
	return nil
}

func (s *Repository) FindOne() error {
	// actual implementation
	return nil
}

func (s *Repository) Update() error {
	// actual implementation
	return nil
}

func (s *Repository) Remove() error {
	// actual implementation
	return nil
}