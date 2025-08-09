package db

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
	SSLMode  string
}

func (c DBConfig) CreateDSN() string {
	user := url.QueryEscape(c.User)
	pass := url.QueryEscape(c.Password)

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		user,
		pass,
		c.Host,
		c.Port,
		c.DBName,
		c.SSLMode,
	)
}

func NewDBConnection(connStr string) (*pgxpool.Pool , error) {
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil{
		log.Fatal(err)	
	}
	return pool , nil
} 
