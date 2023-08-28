package config

import (
	"fmt"
	"os"
)

// Config declares services configuration
type Config struct {
	BindAddr string
}

// Read reads configuration from environment.
func Read() (*Config, error) {
	var (
		config  Config
		isFound bool
	)

	config.BindAddr, isFound = os.LookupEnv("BIND_ADDR")
	if !isFound {
		return nil, fmt.Errorf("cannot find 'BIND_ADDR' in env")
	}

	return &config, nil
}
