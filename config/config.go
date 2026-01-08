package config

import (
	"guardpost-gin/internal/models"
	"os"

	"github.com/goccy/go-yaml"
)

func LoadConfig(filename string) (*models.Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config models.Config
	err = yaml.Unmarshal(data, &config)
	return &config, err
}
