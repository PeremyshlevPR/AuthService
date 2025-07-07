package app

import (
	"os"

	"github.com/PeremyshlevPR/AuthService/internal/db"
	"github.com/PeremyshlevPR/AuthService/internal/grpc"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DB         db.PostgresConfig
	GRPCServer grpc.GRPCServerConfig
}

func LoadConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
