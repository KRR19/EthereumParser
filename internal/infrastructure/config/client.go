package config

import "time"

type Config struct {
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) BlockCheckInterval() time.Duration {
	return 5 * time.Second
}
