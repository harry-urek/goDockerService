package main

import (
	"fmt"
	"os"
)

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
	JWTSecret  string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DBUser", "root"),
		DBPassword: getEnv("DBPassword", "myPassword"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:     getEnv("DB_NAME", "go-rest-service"),
		JWTSecret:  getEnv("JWT_SECRET", "random-jwt-string-key"),
	}

}
func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
