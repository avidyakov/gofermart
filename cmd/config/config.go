package config

import (
	"flag"
	"github.com/caarlos0/env"
	"log"
	"time"
)

type Config struct {
	RunAddress     string        `env:"RUN_ADDRESS"`
	DatabaseURI    string        `env:"DATABASE_URI"`
	SecretKey      string        `env:"SECRET_KEY"`
	TokenExp       time.Duration `env:"TOKEN_EXP"`
	AccrualAddress string        `env:"ACCRUAL_SYSTEM_ADDRESS"`
}

func NewConfig() *Config {
	config := &Config{}
	config.loadConfigFromArgs()
	config.loadConfigFromEnv()
	return config
}

func (c *Config) loadConfigFromArgs() {
	flag.StringVar(&c.RunAddress, "a", ":8000", "run address")
	flag.StringVar(&c.DatabaseURI, "d", "postgres://postgres:changeme@localhost:5432/gofermart", "database uri")
	flag.StringVar(&c.SecretKey, "s", "", "secret key")
	flag.DurationVar(&c.TokenExp, "t", time.Hour*3, "secret key")
	flag.StringVar(&c.AccrualAddress, "r", "http://localhost:8080/api/orders", "accrual system address")
	flag.Parse()
}

func (c *Config) loadConfigFromEnv() {
	if err := env.Parse(c); err != nil {
		log.Fatalf("Failed to parse environment: %v", err)
	}
}
