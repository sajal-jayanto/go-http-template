package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sajal-jayanto/go-http-template/cmd/api"
	"github.com/sajal-jayanto/go-http-template/configs"
	"github.com/sajal-jayanto/go-http-template/db"
)

func main(){
	DBconfig := db.DBConfig{
		User:      configs.Envs.DBUser,
		Password:  configs.Envs.DBPassword,
		Host:      configs.Envs.DBHost,
		Port:      configs.Envs.DBPort,             
		DBName:    configs.Envs.DBName,
		SSLMode:   configs.Envs.DB_SSLMode,
	}
	
	connStr := DBconfig.CreateDSN()

	db, err := db.NewDBConnection(connStr)
	if err != nil {
		log.Fatal(err)
	}

	DBInit(db)

	port := ":" + configs.Envs.Port
	server := api.NewAPIServer(port, db)
	if err := server.Run(); err != nil{
		log.Fatal(err)
	}
}

func DBInit(db *pgxpool.Pool){
	err := db.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}
	log.Println("DB: Successfully connected!")
}