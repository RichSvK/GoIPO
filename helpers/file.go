package helpers

import (
	"os"
	"path/filepath"
)

func ReadFolder(folderName string) ([]string, error) {
	var fileList []string
	err := filepath.Walk(folderName, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Add the file path to the slice
		if !info.IsDir() {
			fileList = append(fileList, path)
		}
		return nil
	})
	return fileList, err
}
