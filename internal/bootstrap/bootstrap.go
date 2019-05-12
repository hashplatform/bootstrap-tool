package bootstrap

import (
	"archive/zip"
	"io"
	"os"
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
	newZip, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newZip.Close()

	zipWriter := zip.NewWriter(newZip)
	defer zipWriter.Close()

	// add files into zip
	for _, file := range files {
		if err = AddFileToZip(zipWriter, file); err != nil {
			return err
		}
	}

	return nil
}

func AddFileToZip(zipWriter *zip.Writer, filename string) error {
	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = filename
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err

}
