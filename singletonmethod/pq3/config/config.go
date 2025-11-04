package config

import (
	"fmt"
	"sync"
)

type Config struct{}

var config *Config
var once sync.Once

func GetConfigInstace() *Config {
	once.Do(func() {
		fmt.Println("Creating a config instance...")
		config = &Config{}
	})
	return config
}

func (*Config) Get(key string) string {
	return fmt.Sprintf("Fetching config value for key: %s", key)
}
