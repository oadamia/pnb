package main

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	OpenAIKey string `envconfig:"openai_api_key" required:"true"`
}

func main() {
	var config Config
	err := envconfig.Process("pnb", &config)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(config.OpenAIKey)
}
