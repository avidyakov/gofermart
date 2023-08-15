package config

import (
	"flag"
	"github.com/caarlos0/env"
	"gophermart/cmd/repo"
	"log"
)

type Config struct {
	RunAddress  string `env:"RUN_ADDRESS"`
	DatabaseURI string `env:"DATABASE_URI"`
	Repo        repo.Repo
}

func NewConfig() *Config {
	config := &Config{}
	config.loadConfigFromArgs()
	config.loadConfigFromEnv()
	//config.Repo = postgres.NewPostgresRepo(config)
	return config
}

func (c *Config) loadConfigFromArgs() {
	flag.StringVar(&c.RunAddress, "a", ":8080", "run address")
	flag.StringVar(&c.DatabaseURI, "d", "postgres://postgres:changeme@localhost:5432/gofermart", "database uri")
	flag.Parse()
}

func (c *Config) loadConfigFromEnv() {
	if err := env.Parse(c); err != nil {
		log.Fatalf("Failed to parse environment: %v", err)
	}
}
