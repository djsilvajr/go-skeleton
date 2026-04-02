package config

import (
	"os"
	"strconv"
)

type Config struct {
	AppName  string
	AppPort  string
	AppEnv   string
	AppDebug bool

	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string

	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int

	JWTSecret     string
	JWTExpireHour int

	MailHost     string
	MailPort     string
	MailUser     string
	MailPassword string
	MailFrom     string

	OtelEnabled          bool
	OtelServiceName      string
	OtelExporterEndpoint string
}

func Load() *Config {
	return &Config{
		AppName:  getEnv("APP_NAME", "go-skeleton"),
		AppPort:  getEnv("APP_PORT", "8020"),
		AppEnv:   getEnv("APP_ENV", "local"),
		AppDebug: getEnvBool("APP_DEBUG", true),

		DBHost:     getEnv("DB_HOST", "db"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBName:     getEnv("DB_DATABASE", "go_skeleton"),
		DBUser:     getEnv("DB_USERNAME", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),

		RedisHost:     getEnv("REDIS_HOST", "redis"),
		RedisPort:     getEnv("REDIS_PORT", "6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", "bananinha123"),
		RedisDB:       getEnvInt("REDIS_DB", 0),

		JWTSecret:     getEnv("JWT_SECRET", "change-me-in-production"),
		JWTExpireHour: getEnvInt("JWT_EXPIRE_HOUR", 24),

		MailHost:     getEnv("MAIL_HOST", "smtp.gmail.com"),
		MailPort:     getEnv("MAIL_PORT", "587"),
		MailUser:     getEnv("MAIL_USERNAME", ""),
		MailPassword: getEnv("MAIL_PASSWORD", ""),
		MailFrom:     getEnv("MAIL_FROM_ADDRESS", "noreply@example.com"),

		OtelEnabled:          getEnvBool("OTEL_ENABLED", false),
		OtelServiceName:      getEnv("OTEL_SERVICE_NAME", "go-skeleton"),
		OtelExporterEndpoint: getEnv("OTEL_EXPORTER_OTLP_ENDPOINT", "http://otel-collector:4318"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvBool(key string, fallback bool) bool {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		return fallback
	}
	return b
}

func getEnvInt(key string, fallback int) int {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		return fallback
	}
	return i
}
