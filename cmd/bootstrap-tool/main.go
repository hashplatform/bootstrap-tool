package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"path/filepath"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var currentDate = time.Now().Format(time.RFC3339)

type Config struct {
	Coin string `json:"coin"`
	Directory string `json:"directory"`
	Destination string `json:"destination"`
}

func LoadConfigFile(file string) Config {
	var config Config
	configurationFile, err := os.Open(file)

	defer configurationFile.Close()

	if err != nil {
		fmt.Println(err.Error())
	}

	json.NewDecoder(configurationFile).Decode(&config)

	return config
}

func ListFiles(config Config) (name string) {

	coin := config.Coin

	bootstrapName := strings.ToLower(coin) + "-" + currentDate

	fmt.Println("=========================")
	fmt.Println("Coin Name:", coin)
	fmt.Println("Blockchain Bootstrap Date:", currentDate)
	fmt.Println("Bootstrap Name:", bootstrapName)
	fmt.Println("=========================")

	files, err := ioutil.ReadDir(config.Directory)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	return bootstrapName
}

func Zip(pathToZip, destinationPath string) error {

	destinationFile, err := os.Create(destinationPath)

	if err != nil {
		return err
	}

	myZip := zip.NewWriter(destinationFile)
	err = filepath.Walk(pathToZip, func(filePath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if err != nil {
			return err
		}

		relPath := strings.TrimPrefix(filePath, filepath.Dir(pathToZip))
		zipFile, err := myZip.Create(relPath)
		if err != nil {
			return err
		}
		fsFile, err := os.Open(filePath)
		if err != nil {
			return err
		}
		_, err = io.Copy(zipFile, fsFile)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	err = myZip.Close()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	config := LoadConfigFile("config.json")

	bootstrapName := strings.ToLower(config.Coin) + "-" + currentDate

	destination := config.Destination + bootstrapName

	ListFiles(config)
	Zip(config.Directory, destination)
}