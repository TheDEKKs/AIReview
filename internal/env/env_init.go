package env

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	KeyAPI string `env:"TOKEN_API,notEmpty"`
}


func (s *Config) Load() {
	godotenv.Load()
	if err := env.Parse(s); err != nil {
		log.Fatalf("couldn't load config: %s", err.Error())
	}
}
