package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

func main() {
	var config Config
	err := envconfig.Process("pnb", &config)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(config.OpenAIAPIKey)
}
