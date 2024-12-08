package config

import (
	"runtime"
	"testing"
	"time"
)

func TestNewConfig(t *testing.T) {
	cfg := NewConfig()

	if cfg.BlockCheckInterval() != 5*time.Second {
		t.Errorf("Expected block check interval to be 5 seconds, got %v", cfg.BlockCheckInterval())
	}

	if cfg.CoreCount() != runtime.NumCPU() {
		t.Errorf("Expected core count to be %d, got %d", runtime.NumCPU(), cfg.CoreCount())
	}
}

func TestBlockCheckInterval(t *testing.T) {
	cfg := NewConfig()
	expected := 5 * time.Second

	if cfg.BlockCheckInterval() != expected {
		t.Errorf("Expected block check interval to be %v, got %v", expected, cfg.BlockCheckInterval())
	}
}

func TestCoreCount(t *testing.T) {
	cfg := NewConfig()
	expected := runtime.NumCPU()

	if cfg.CoreCount() != expected {
		t.Errorf("Expected core count to be %d, got %d", expected, cfg.CoreCount())
	}
}
