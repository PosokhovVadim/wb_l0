package config

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ServiceName string     `yaml:"service_name" env-required:"true"`
	Env         string     `yaml:"env" env-required:"true"`
	StoragePath string     `yaml:"storage_path" env-required:"true"`
	RedisPath   string     `yaml:"redis_path" env-required:"true"`
	NatsUrl     string     `yaml:"nats_url" env-required:"true"`
	HTTPServer  HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Port    int           `yaml:"port" env-required:"true"`
	Timeout time.Duration `yaml:"timeout" env-required:"true"`
}

func New(filePath string) (*Config, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return &Config{}, err
	}
	var config Config
	err := cleanenv.ReadConfig(filePath, &config)

	return &config, err
}
