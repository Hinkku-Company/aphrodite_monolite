package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Hinkku-Company/aphrodite_monolite/logger"
)

type redis struct {
	RedisHost     string
	RedisPort     string
	RedisDB       string
	RedisPassword string
}

type dataBase struct {
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBIsInsecure bool
}

type jwt struct {
	AccessTokenPrivateKey  string
	AccessTokenPublicKey   string
	AccessTokenExpiredMin  string
	RefreshTokenPrivateKey string
	RefreshTokenPublicKey  string
	RefreshTokenExpiredMin string
}

type api struct {
	APPPort string
}

type Config struct {
	AppENV        string
	AdminPassword string
	dataBase
	api
	redis
	jwt
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) LoadConfigFromEnv() (Config, error) {
	logger.Log().Info("Load .env configuration")
	var config Config
	requiredEnvVars := []string{
		"ADMIN_PASSWORD",
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
		"REDIS_HOST",
		"REDIS_PORT",
		"REDIS_PASSWORD",
		"ACCESS_TOKEN_PRIVATE_KEY",
		"ACCESS_TOKEN_PUBLIC_KEY",
		"ACCESS_TOKEN_EXPIRED_MIN",
		"REFRESH_TOKEN_PRIVATE_KEY",
		"REFRESH_TOKEN_PUBLIC_KEY",
		"REFRESH_TOKEN_EXPIRED_MIN",
	}
	if err := c.checkRequiredEnvVars(requiredEnvVars); err != nil {
		return config, err
	}

	config = Config{
		AppENV:        c.getEnv("APP_ENV"),
		AdminPassword: c.getEnv("ADMIN_PASSWORD"),
		dataBase: dataBase{
			DBHost:       c.getEnv("DB_HOST"),
			DBPort:       c.getEnv("DB_PORT"),
			DBUser:       c.getEnv("DB_USER"),
			DBPassword:   c.getEnv("DB_PASSWORD"),
			DBName:       c.getEnv("DB_NAME"),
			DBIsInsecure: c.getBoolEnv("DB_INSECURE", false),
		},
		api: api{
			APPPort: c.getEnvDefault("APP_PORT", "8118"),
		},
		redis: redis{
			RedisHost:     c.getEnv("REDIS_HOST"),
			RedisPort:     c.getEnv("REDIS_PORT"),
			RedisDB:       c.getEnvDefault("REDIS_DB", "0"),
			RedisPassword: c.getEnv("REDIS_PASSWORD"),
		},
		jwt: jwt{
			AccessTokenPrivateKey:  c.getEnv("ACCESS_TOKEN_PRIVATE_KEY"),
			AccessTokenPublicKey:   c.getEnv("ACCESS_TOKEN_PUBLIC_KEY"),
			AccessTokenExpiredMin:  c.getEnv("ACCESS_TOKEN_EXPIRED_MIN"),
			RefreshTokenPrivateKey: c.getEnv("REFRESH_TOKEN_PRIVATE_KEY"),
			RefreshTokenPublicKey:  c.getEnv("REFRESH_TOKEN_PUBLIC_KEY"),
			RefreshTokenExpiredMin: c.getEnv("REFRESH_TOKEN_EXPIRED_MIN"),
		},
	}

	return config, nil
}

func (c *Config) checkRequiredEnvVars(requiredVars []string) error {
	for _, envVar := range requiredVars {
		if c.getEnv(envVar) == "" {
			return fmt.Errorf("la variable de entorno %s no está definida", envVar)
		}
	}
	return nil
}

func (c *Config) getEnvDefault(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func (c *Config) getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return ""
	}
	return value
}

func (c *Config) getBoolEnv(key string, defaultValue bool) bool {
	valueStr := c.getEnv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}
