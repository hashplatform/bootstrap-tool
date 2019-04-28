package bootstrap

import (
	"log"
	"os"
	"path/filepath"
)

func ListBootstrapFiles(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal("An error occurred:", err)
		}

		*files = append(*files, path)
		return nil
	}
}
