package config

import (
	"os"

	"github.com/go-viper/mapstructure/v2"
	"github.com/pelletier/go-toml"
)

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Unmarshal into a generic map to avoid losing unmapped fields,
	// which exist because integrations can have arbitrary fields
	var rawConfig map[string]interface{}
	if err := toml.Unmarshal(data, &rawConfig); err != nil {
		return nil, err
	}

	// Decode into a Config struct, with unmapped fields captured by
	// struct fields with the tag `mapstructure:",remain"`
	var cfg Config
	if err := mapstructure.Decode(rawConfig, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
