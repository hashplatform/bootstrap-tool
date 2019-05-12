package bootstrap

import (
	"archive/zip"
	"io"
	"os"
)

func AppendFiles(filename string, zipw *zip.Writer) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	wr, err := zipw.Create(filename)
	if err != nil {
		return err
	}

	if _, err := io.Copy(wr, file); err != nil {
		return err
	}

	return nil
}
