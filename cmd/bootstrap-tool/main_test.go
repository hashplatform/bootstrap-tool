package main

import (
	"os"
	"testing"
)

func TestLoadConfigFile(t *testing.T) {
	config := struct {
		Coin        string
		Directory   string
	}{
		"bitcoin",
		"../../testData",
	}

	if config.Coin == "" {
		t.Errorf("Coin parameter is empty, got: %s, want: %s", config.Coin, "bitcoin")
	}
	if config.Directory == "" {
		t.Errorf("Directory parameter is empty, got: %s, want: %s", config.Directory, "../../testData")
	}
}

func TestListFiles(t *testing.T) {
	if _, err := os.Stat("../../testData"); err != nil {
		if os.IsNotExist(err) {
			t.Errorf("Directory does not exist")
		}
	}
}
