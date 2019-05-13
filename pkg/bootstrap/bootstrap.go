package bootstrap

import (
	"github.com/mholt/archiver"
)

// check for important files
func CheckFileSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

// zipping process
func ZipBlockchain(filename string, files []string) error {
	err := archiver.Archive(files, filename)
	if err != nil {
		return err
	}

	return nil
}
