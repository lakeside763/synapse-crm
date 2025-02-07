package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	DatabaseUrl  string
	RedisUrl     string
	DatabaseName string
	Port         int
}

func NewConfig() *Config {
	databaseUrl := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	port, _ := strconv.Atoi(os.Getenv("PORT"))

	return &Config{
		DatabaseUrl:  databaseUrl,
		RedisUrl:     os.Getenv("REDIS_URL"),
		DatabaseName: os.Getenv("DB_NAME"),
		Port:         port,
	}
}
