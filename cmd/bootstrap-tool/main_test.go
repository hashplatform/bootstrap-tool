package main

import "testing"

func TestLoadConfigFile(t *testing.T) {
	config := struct {
		Coin string
		Directory string
		Destination string
	} {
		"bitcoin",
		"../../testData",
		"../../testingArchives",
	}

	if config.Coin == "" {
		t.Errorf("Coin paramter is empty, got: %s, want: %s", config.Coin, "bitcoin")
	}
	if config.Directory == "" {
		t.Errorf("Directory paramter is empty, got: %s, want: %s", config.Directory, "../../testData")
	}
	if config.Destination == "" {
		t.Errorf("Directory paramter is empty, got: %s, want: %s", config.Destination, "../../testingArchives")
	}
}