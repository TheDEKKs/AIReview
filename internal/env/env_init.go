package env

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	KeyAPI string `env:"TOKEN_API,notEmpty"`
}


func (s *Config) Load() {
	ex, _ := os.Executable()
    dir := filepath.Dir(ex)
    
    godotenv.Load(filepath.Join(dir, ".env"))
    
	if err := env.Parse(s); err != nil {
		log.Fatalf("couldn't load config: %s", err.Error())
	}
}
