package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sajal-jayanto/go-http-template/src/sample"
)

type APIServer struct {
	addr string
	db *pgxpool.Pool
}

func NewAPIServer(addr string, db *pgxpool.Pool) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (server *APIServer) Run() error {
	mainRouter := gin.Default()
	api := mainRouter.Group("/api")
	
	v1 := api.Group("/v1")
	{
		// sample module 
		routerGroup := v1.Group("/sample")
		sampleRepo := sample.NewRepository(server.db)
		sampleHandler := sample.NewHandler(sampleRepo)
		sampleHandler.RegisterRoutes(routerGroup)

	}

	log.Println("Server started on port", server.addr)
	return http.ListenAndServe(server.addr, mainRouter)
} 

