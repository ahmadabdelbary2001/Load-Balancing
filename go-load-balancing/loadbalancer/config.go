package loadbalancer

import (
	"encoding/json"
	"os"
)

type Config struct {
	ListenPort          string   `json:"listenPort"`
	Servers             []string `json:"servers"`
	HealthCheckInterval string   `json:"healthCheckInterval"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
