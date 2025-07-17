package config

import (
	"os"
	"testing"
)

func TestLoadDefaults(t *testing.T) {
	os.Clearenv()
	cfg := Load()

	if cfg.DBHost != "postgres" {
		t.Errorf("expected default DBHost postgres, got %s", cfg.DBHost)
	}
	if cfg.DBPort != "5432" {
		t.Errorf("expected default DBPort 5432, got %s", cfg.DBPort)
	}
	if cfg.DBUser != "postgres" {
		t.Errorf("expected default DBUser postgres, got %s", cfg.DBUser)
	}
	if cfg.DBPassword != "postgres" {
		t.Errorf("expected default DBPassword postgres, got %s", cfg.DBPassword)
	}
	if cfg.DBName != "app" {
		t.Errorf("expected default DBName app, got %s", cfg.DBName)
	}
	if cfg.ServerPort != "8080" {
		t.Errorf("expected default ServerPort 8080, got %s", cfg.ServerPort)
	}
}

func TestLoadFromEnv(t *testing.T) {
	os.Clearenv()
	os.Setenv("DB_HOST", "h")
	os.Setenv("DB_PORT", "p")
	os.Setenv("DB_USER", "u")
	os.Setenv("DB_PASSWORD", "pass")
	os.Setenv("DB_NAME", "db")
	os.Setenv("SERVER_PORT", "9000")

	cfg := Load()

	if cfg.DBHost != "h" || cfg.DBPort != "p" || cfg.DBUser != "u" || cfg.DBPassword != "pass" || cfg.DBName != "db" || cfg.ServerPort != "9000" {
		t.Errorf("env variables not loaded correctly: %+v", cfg)
	}
}
