package config

import "testing"

func TestLoadDefaults(t *testing.T) {
	t.Setenv("ADDR", "")
	t.Setenv("SHUTDOWN_TIMEOUT_SEC", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cfg.Addr != ":8080" {
		t.Fatalf("unexpected addr: %s", cfg.Addr)
	}
	if cfg.ShutdownTimeout.Seconds() != 10 {
		t.Fatalf("unexpected timeout: %v", cfg.ShutdownTimeout)
	}
}

func TestLoadCustom(t *testing.T) {
	t.Setenv("ADDR", ":9090")
	t.Setenv("SHUTDOWN_TIMEOUT_SEC", "5")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cfg.Addr != ":9090" {
		t.Fatalf("unexpected addr: %s", cfg.Addr)
	}
	if cfg.ShutdownTimeout.Seconds() != 5 {
		t.Fatalf("unexpected timeout: %v", cfg.ShutdownTimeout)
	}
}
