package main

import (
	"errors"
	"os"
	"strings"
)

func scandir(dirPath string) ([]string, error) {
	f, err := os.Open(dirPath)
	if err != nil {
		return []string{}, errors.New("Unable to open directory")
	}
	// Open the directory and return an error if it fails.

	files, err := f.Readdir(0)
	if err != nil {
		return []string{}, errors.New("Unable to open directory")
	}
	// Read the contents of the directory and return an error if it fails.

	result := make([]string, len(files))
	for i := range files {
		result[i] = strings.Split(files[i].Name(), ".")[0]
		// Extract the file name without the extension.
	}
	// Loop through the files and extract the file name without the extension.

	return result, nil
}

// scandir is a function that reads the contents of a directory and returns a list of file names.
