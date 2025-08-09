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
	Create(*gin.Context, types.CreateSampleReq) (*types.Sample, error)
	FindAll(*gin.Context) ([]*types.Sample, error) 
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

func (repository *Repository) FindAll(ctx *gin.Context) ([]*types.Sample, error) {
	rows, err := repository.db.Query(ctx ,"SELECT id, data FROM sample")
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	var samples []*types.Sample
	for rows.Next() {
		var sample types.Sample
		if err := rows.Scan(&sample.Id, &sample.Data); err != nil {
			return nil, err
		}
		samples = append(samples, &sample)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return samples, nil
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