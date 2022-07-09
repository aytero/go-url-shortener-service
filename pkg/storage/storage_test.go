package storage

import (
	"testing"
	"urlShortener/pkg/config"
)

func TestLocalStorageInit(t *testing.T) {
	_, err := NewStorage(&config.Config{StorageType: "lol"})
	if err == nil {
		t.Fatalf("Inited unknown storage type\n")
	}

	_, err = NewStorage(&config.Config{StorageType: "local"})
	if err != nil {
		t.Fatalf("Failed to init local storage\n")
	}

	_, err = NewStorage(&config.Config{StorageType: "postgres"})
	if err != nil {
		t.Fatalf("Failed to init postgres storage\n")
	}
}
