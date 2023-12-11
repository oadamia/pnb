package main

import (
	"fmt"
	"log/slog"

	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	OpenAIKey string `envconfig:"openai_api_key" required:"true"`
}

func main() {
	slog.Info("Starting up...")

	var cfg Config
	err := envconfig.Process("pnb", &cfg)
	if err != nil {
		fmt.Println(err)
	}

}
