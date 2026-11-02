package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config 保存服务运行所需的配置。
type Config struct {
	Addr            string
	ShutdownTimeout time.Duration
}

// Load 从环境变量读取配置，未设置时使用合理默认值。
func Load() (Config, error) {
	addr := getenv("ADDR", ":8080")
	timeoutSec, err := strconv.Atoi(getenv("SHUTDOWN_TIMEOUT_SEC", "10"))
	if err != nil {
		return Config{}, fmt.Errorf("invalid SHUTDOWN_TIMEOUT_SEC: %w", err)
	}
	if timeoutSec <= 0 {
		return Config{}, fmt.Errorf("SHUTDOWN_TIMEOUT_SEC must be positive")
	}

	return Config{
		Addr:            addr,
		ShutdownTimeout: time.Duration(timeoutSec) * time.Second,
	}, nil
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
