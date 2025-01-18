package config

import (
	"fmt"

	"github.com/VolodymyrShabat/TestMicroservices/resource-service/internal/models"
	"github.com/spf13/viper"
)

func LoadConfig(path string) (*models.Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg models.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}
