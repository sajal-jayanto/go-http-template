package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	pgxMigrate "github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sajal-jayanto/go-http-template/configs"
	"github.com/sajal-jayanto/go-http-template/db"
)

func main() {
	DBconfig := db.DBConfig{
		User:      configs.Envs.DBUser,
		Password:  configs.Envs.DBPassword,
		Host:      configs.Envs.DBHost,
		Port:      configs.Envs.DBPort,             
		DBName:    configs.Envs.DBName,
		SSLMode:   configs.Envs.DB_SSLMode,
	}
	
	connStr := DBconfig.CreateDSN()
	db, err := sql.Open("postgres", connStr)

	driver, err := pgxMigrate.WithInstance(db, &pgxMigrate.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "pgx", driver)
	if err != nil {
		log.Fatal(err)
	}

	v, d, _ := m.Version()
	log.Printf("Version: %d, dirty: %v", v, d)

	cmd := os.Args[len(os.Args) - 1]
	
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}
