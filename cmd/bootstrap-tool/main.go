package main

import (
	"archive/zip"
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

	json.NewDecoder(configurationFile).Decode(&config)

	return config
}

func ZipWriter(coin Config) {

	directory := coin.Directory

	// Create a buffer
	outputFile, err := os.Create(directory)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer outputFile.Close()

	// Create archive
	w := zip.NewWriter(outputFile)

	//add files
	addFiles(w, directory, "")

	if err != nil {
		fmt.Println(err.Error())
	}

	// Close file
	err = w.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func addFiles(w *zip.Writer, basePath, baseInZip string) {
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, file := range files {
		fmt.Println(basePath + file.Name())
		if !file.IsDir() {
			dat, err := ioutil.ReadFile(basePath + file.Name())
			if err != nil {
				fmt.Println(err.Error())
			}

			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				fmt.Println(err.Error())
			}
			_, err = f.Write(dat)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else if file.IsDir() {
			newBase := basePath + file.Name() + "/"
			fmt.Println("Recursing and Adding SubDir: " + file.Name())
			fmt.Println("Recursing and Adding SubDir: " + newBase)

			addFiles(w, newBase, file.Name() + "/")
		}
	}
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
	ZipWriter(config)
}