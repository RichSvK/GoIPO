package configs

import (
	"fmt"
	"os"
)

func MakeOutputFolder(folderName string) {
	// Check if there is a "output" directory in current directory
	path := "./" + folderName
	_, checkFolder := os.Stat(path)

	// If checkFolder != nil means there is no "output" directory in current directory
	if checkFolder != nil {
		// Make "output" directory
		for checkFolder != nil {
			err := os.Mkdir(path, 0755)
			if err != nil {
				fmt.Println("Error creating directory")
			}
			_, checkFolder = os.Stat(path)
		}
	}
}
