package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	PublicHost              string
	Port                    string
	DBUser                  string
	DBPassword              string
	DBSource                string
	DBName                  string
	CookiesAuthSecret       string
	CookiesAuthAgeInSeconds int
	CookiesAuthIsSecure     bool
	CookiesAuthIsHttpOnly   bool
	DiscordClientID         string
	DiscordClientSecret     string
	GithubClientID          string
	GithubClientSecret      string
	SymmetricKey            string
}

const (
	twoDaysInSeconds = 60 * 60 * 24 * 2
)

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	return initConfig()
}

var Envs = LoadConfig()

func initConfig() Config {
	return Config{
		PublicHost:              getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                    getEnv("PORT", "8080"),
		DBUser:                  getEnv("DB_USER", "root"),
		DBPassword:              getEnv("DB_PASSWORD", "mypassword"),
		DBSource:                getEnvOrError("DB_SOURCE"),
		SymmetricKey:            getEnvOrError("SYMMETRIC_KEY"),
		DBName:                  getEnv("DB_NAME", "cars"),
		CookiesAuthSecret:       getEnv("COOKIES_AUTH_SECRET", "some-very-secret-key"),
		CookiesAuthAgeInSeconds: getEnvAsInt("COOKIES_AUTH_AGE_IN_SECONDS", twoDaysInSeconds),
		CookiesAuthIsSecure:     getEnvAsBool("COOKIES_AUTH_IS_SECURE", false),
		CookiesAuthIsHttpOnly:   getEnvAsBool("COOKIES_AUTH_IS_HTTP_ONLY", false),
		DiscordClientID:         getEnvOrError("DISCORD_CLIENT_ID"),
		DiscordClientSecret:     getEnvOrError("DISCORD_CLIENT_SECRET"),
		GithubClientID:          getEnvOrError("GITHUB_CLIENT_ID"),
		GithubClientSecret:      getEnvOrError("GITHUB_CLIENT_SECRET"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvOrError(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	panic(fmt.Sprintf("Environment variable %s is not set", key))

}

func getEnvAsInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}

func getEnvAsBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		b, err := strconv.ParseBool(value)
		if err != nil {
			return fallback
		}

		return b
	}

	return fallback
}
