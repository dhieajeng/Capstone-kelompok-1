package configs

import (
	"errors"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Env       string         `env:"ENV" envDefault:"dev"`
	URL       string         `env:"URL" envDefault:"http://localhost"`
	Port      string         `env:"PORT" envDefault:"8080"`
	Postgres  PostgresConfig `envPrefix:"POSTGRES_"`
	JWT       JwtConfig      `envPrefix:"JWT_"`
	Redis     RedisConfig    `envPrefix:"REDIS_"`
	Namespace NamespaceConfig
	SMTP      SMTPConfig     `envPrefix:"SMTP_"`
	Midtrans  MidtransConfig `envPrefix:"MIDTRANS_"`
}

type PostgresConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     string `env:"PORT" envDefault:"5432"`
	User     string `env:"USER" envDefault:"postgres"`
	Password string `env:"PASSWORD" envDefault:"postgres"`
	Database string `env:"DATABASE" envDefault:"postgres"`
}

type JwtConfig struct {
	SecretKey string `env:"SECRET_KEY"`
}

type RedisConfig struct {
	Host     string `env:"HOST_REDIS" envDefault:"localhost"`
	Port     string `env:"PORT_REDIS" envDefault:"6379"`
	Password string `env:"PASSWORD_REDIS" envDefault:""`
}

type NamespaceConfig struct {
	Namespace string `env:"NAMESPACE" envDefault:"application_namespace"`
}

type SMTPConfig struct {
	Host     string `env:"HOST" envDefault:""`
	Port     int    `env:"PORT" envDefault:""`
	Username string `env:"USERNAME" envDefault:""`
	Password string `env:"PASSWORD" envDefault:""`
	Sender   string `env:"SENDER" envDefault:""`
}

type MidtransConfig struct {
	ServerKey string `env:"SERVER_KEY" envDefault:""`
	ClientKey string `env:"CLIENT_KEY" envDefault:""`
}

func NewConfig(path string) (*Config, error) {
	err := godotenv.Load(path)
	if err != nil {
		return nil, errors.New("failed to load .env")
	}

	cfg := new(Config)

	err = env.Parse(cfg)
	if err != nil {
		return nil, errors.New("failed to parse config file")
	}

	return cfg, nil
}
