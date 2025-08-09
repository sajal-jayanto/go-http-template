package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	DBUser       string
	DBPassword   string
	DBHost       string 
	DBPort       int                  
	DBName       string
	DB_SSLMode   string 
	Port         string
}

var Envs = initConfig()

func initConfig() EnvConfig {
	godotenv.Load()

	port, err := strconv.Atoi(getEnv("DB_PORT" , "5432"))
	if err != nil{
		port = 5432
	}
	
	return EnvConfig{
		DBUser:      getEnv("DB_USER" , "root"),
		DBPassword:  getEnv("DB_PASSWORD" , "root"),
		DBHost:      getEnv("DB_HOST" , "localhost"),
		DBPort:      port,
		DBName:      getEnv("DB_NAME" , "test"),
		DB_SSLMode:  getEnv("DB_SSL_MODE" , "disable"),
		Port:        getEnv("PORT" , "8080"),
	}
}


func getEnv(key , fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}