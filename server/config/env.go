package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort         string
	DB_DSN          string
	JWTAccessToken  string
	JWTRefreshToken string
	RedisHost       string
	RedisPassword   string
}

func Env() Config {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	app_port, errAppPort := os.LookupEnv("APP_PORT")
	if !errAppPort {
		panic("APP_PORT ENV not found!")
	}

	db, errDb := os.LookupEnv("DB_DSN")
	if !errDb {
		panic("DB_DSN env not found!")
	}

	accessToken, errAccessToken := os.LookupEnv("JWT_ACCESS_TOKEN_SECRET")
	if !errAccessToken {
		panic("JWT_ACCESS_TOKEN_SECRET env not found!")
	}

	refreshToken, errRefreshToken := os.LookupEnv("JWT_REFRESH_TOKEN_SECRET")
	if !errRefreshToken {
		panic("JWT_REFRESH_TOKEN_SECRET env not found!")
	}

	redisHost, errRedisHost := os.LookupEnv("REDIS_HOST")
	if !errRedisHost {
		panic("REDIS_HOST env not found!")
	}

	redisPassword, errRedisPassword := os.LookupEnv("REDIS_PASSWORD")
	if !errRedisPassword {
		panic("REDIS_PASSWORD env not found!")
	}

	result := Config{
		AppPort:         app_port,
		DB_DSN:          db,
		JWTAccessToken:  accessToken,
		JWTRefreshToken: refreshToken,
		RedisHost:       redisHost,
		RedisPassword:   redisPassword,
	}

	return result
}
