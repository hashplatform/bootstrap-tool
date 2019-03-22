package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var currentDate = time.Now().Format(time.RFC3339)

type Config struct {
	Coin string `json:"coin"`
	Directory string `json:"directory"`
}

func LoadConfigFile(file string) Config {
	var config Config
	configurationFile, err := os.Open(file)

	defer configurationFile.Close()

	if err != nil {
		fmt.Println(err.Error())
	}

	jsonParser := json.NewDecoder(configurationFile)
	jsonParser.Decode(&config)
	return config
}

func ListFiles(config Config) {

	coin := config.Coin
	fmt.Println("=========================")
	fmt.Println("Coin Name:", coin)
	fmt.Println("Blockchain Bootstrap Date:", currentDate)
	fmt.Println("Bootstrap Name:", strings.ToLower(coin) + "-" + currentDate)
	fmt.Println("=========================")

	files, err := ioutil.ReadDir(config.Directory)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func main() {
	config := LoadConfigFile("config.json")
	ListFiles(config)
}