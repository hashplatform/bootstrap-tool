package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Coin string `json:"coin"`
	Directory string `json:"directory"`
}

func LoadConfig(path string) (Config, error) {
	var config Config

	configJson, err := os.Open(path)
	defer configJson.Close()
	if err != nil {
		return Config{}, err
	}

	bytes, err := ioutil.ReadAll(configJson)
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

