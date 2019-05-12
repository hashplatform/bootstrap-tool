package main

import (
	"flag"
	"io/ioutil"

	//"github.com/jackkdev/bootstrap-tool/internal/bootstrap"
	"github.com/jackkdev/bootstrap-tool/internal/config"
	"github.com/jackkdev/bootstrap-tool/internal/preamble"
	"log"
	"time"
)

var version = "1.0.0"

var coin string
var directory string

var currentDate = time.Now().Format(time.RFC3339)

func main() {

	preamble.Preamble(version)

	time.Sleep(2 * time.Second)

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

		// load files from directory and append to list
		files, err := ioutil.ReadDir(directory)
		if err != nil {
			log.Fatal("An error occurred reading from the directory:", err)
		}

		var fileList []string

		for _, file := range files {
			fileList = append(fileList, file.Name())
		}

		log.Println(fileList)

		log.Println("Generating bootstrap for", coin, "at the directory:", directory)
	} else {
		log.Fatal(`No configuration file was provided. Please restart the application and specify the configuration file with -config="config.json"`)
	}
}
