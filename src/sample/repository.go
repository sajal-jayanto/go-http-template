package sample

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sajal-jayanto/go-http-template/types"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

type SampleRepo interface {
	Create(*gin.Context , types.CreateSampleReq) (*types.Sample, error)
	FindAll() error
	FindOne() error
	Update()  error
	Remove()  error
}

func (repository *Repository) Create(ctx *gin.Context, sample types.CreateSampleReq) (*types.Sample, error) {
	var createdSample types.Sample
	err := repository.db.QueryRow(ctx, 
		"INSERT INTO sample (data) VALUES ($1) RETURNING id, data", 
		sample.Data,
	).Scan(&createdSample.Id, &createdSample.Data)
	if err != nil {	
		return nil, err
	}
	
	return &createdSample, nil
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