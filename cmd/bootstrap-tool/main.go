package main

import (
	"flag"
	"fmt"
	"github.com/jackkdev/bootstrap-tool/internal/bootstrap"
	"github.com/jackkdev/bootstrap-tool/internal/config"
	"github.com/jackkdev/bootstrap-tool/internal/preamble"
	"log"
	"path/filepath"
	"time"
)

var version = "1.0.0"

var coin string
var directory string

var currentDate = time.Now().Format(time.RFC3339)

func main() {

	preamble.Preamble(version)
	//time.Sleep(10 * time.Second)

	var bootstrapConfig string

	flag.StringVar(&bootstrapConfig, "config", "", "Directory of your configuration file")
	flag.Parse()

	if bootstrapConfig != "" {
		bootstrapConfigInfo, err := config.LoadConfig(bootstrapConfig)
		if err != nil {
			log.Println("Error reading config file from:", bootstrapConfig)
			return
		} else {
			coin = bootstrapConfigInfo.Coin
			directory = bootstrapConfigInfo.Directory
		}

		log.Println("Loaded configuration file successfully!")
		log.Println("Generating bootstrap for", coin, "at the directory:", directory)
		fmt.Println("")

		var files []string

		err = filepath.Walk(directory, bootstrap.ListBootstrapFiles(&files))
		if err != nil {
			log.Fatal("An error occurred:", err)
		}

		for _, file := range files {
			fmt.Println(file)
		}
	}
}
