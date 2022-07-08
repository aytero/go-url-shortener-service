package config

import (
	"os"
	"testing"
)

func TestConfig(t *testing.T) {

	//loadEnvFile(".env")
	cfg := NewConfig("../../.env")

	want := os.Getenv("URL_HOST")
	if cfg.UrlHost != want {
		t.Errorf("Expected: %s; got: %s\n", want, cfg.UrlHost)
	}
	want = os.Getenv("DATABASE_URL")
	if cfg.DatabaseUrl != want {
		t.Errorf("Expected: %s; got: %s\n", want, cfg.DatabaseUrl)
	}
}

func TestConfigPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	_ = NewConfig(".wrongEnv")
}
