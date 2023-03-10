package main

import (
	"fmt"

	"github.com/blackpointcyber/configparser"
)

func main() {
	var config struct {
		SecretKey int    `yaml:"secretKey" validate:"required"`
		BaseURL   string `yaml:"baseUrl" default:"https://example.com"`

		Address struct {
			Street  string `yaml:"street" default:"defaultStreet"`
			City    string `yaml:"city"`
			Country string `yaml:"country"`
		} `yaml:"address"`
	}

	configparser.MustParseYAMLFile("./examples/simple_usage/config.yaml", &config)

	fmt.Println(config)
}
