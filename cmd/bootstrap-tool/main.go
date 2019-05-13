package main

import (
	"flag"
	"io/ioutil"

	"github.com/jackkdev/bootstrap-tool/pkg/bootstrap"
	"github.com/jackkdev/bootstrap-tool/pkg/config"
	"github.com/jackkdev/bootstrap-tool/pkg/preamble"
	"log"
	"time"
)

var version = "1.0.1"

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

		name := coin + "-" + currentDate + ".zip"

		log.Println("Loaded configuration file successfully!")

		// load files from directory and append to list
		files, err := ioutil.ReadDir(directory)
		if err != nil {
			log.Fatal("An error occurred reading from the directory:", err)
		}

		var fileList []string

		for _, file := range files {
			fileList = append(fileList, directory+file.Name())

			// remove wallet.dat
			if bootstrap.CheckFileSlice(directory+"wallet.dat", fileList) {
				fileList = fileList[:len(fileList)-1]
			}
			if bootstrap.CheckFileSlice(directory+"backups", fileList) {
				fileList = fileList[:len(fileList)-1]
			}
			if bootstrap.CheckFileSlice(directory+"debug.log", fileList) {
				fileList = fileList[:len(fileList)-1]
			}
			if bootstrap.CheckFileSlice(directory+"masternode.conf", fileList) {
				fileList = fileList[:len(fileList)-1]
			}
			if bootstrap.CheckFileSlice(directory+coin+".conf", fileList) {
				fileList = fileList[:len(fileList)-1]
			}
			if bootstrap.CheckFileSlice(directory+"db.log", fileList) {
				fileList = fileList[:len(fileList)-1]
			}
		}

		log.Println("Fetching blockchain files for", coin, "located at the directory:", directory)
		log.Println("Creating bootstrap, please sit tight. The bootstrap name is:", name)

		// start bootstrapping process
		if err = bootstrap.ZipBlockchain(name, fileList); err != nil {
			log.Fatal(err)
		} else if err == nil {
			log.Println("Bootstrapping complete. It can be found in your current directory.")
		}

	} else {
		log.Fatal(`No configuration file was provided. Please restart the application and specify the configuration file with -config="config.json"`)
	}

	log.Println("Thank you for using the Blockchain Bootstrapper :)")
}
