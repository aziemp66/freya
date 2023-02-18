package env

import (
	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	Port         int    `env:"PORT,unset" envDefault:"5000"`
	JwtSecretKey string `env:"JWT_SECRET_KEY,unset"`
	GinMode      string `env:"GIN_MODE,unset" envDefault:"debug"`
}

func LoadConfig() *config {
	cfg := new(config)
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}
