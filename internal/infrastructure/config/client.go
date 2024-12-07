package config

import (
	"runtime"
	"time"
)

type Config struct {
	blockCheckInterval time.Duration
	coreCount          int
}

func NewConfig() *Config {
	return &Config{
		blockCheckInterval: 5 * time.Second,
		coreCount:          runtime.NumCPU(),
	}
}

func (c *Config) BlockCheckInterval() time.Duration {
	return c.blockCheckInterval
}

func (c *Config) CoreCount() int {
	return c.coreCount
}
