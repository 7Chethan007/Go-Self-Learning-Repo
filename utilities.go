package main

import (
	"errors"
	"os"
	"strings"
)

// The os dir package is used to read the contents of a directory.

func scandir(dirPath string) ([]string, error) {
	f, err := os.Open(dirPath)
	if err != nil {
		return []string{}, errors.New("Unable to access directory " + dirPath)
	}
	// Open the directory and return an error if it fails.

	files, err := f.Readdir(0)
	if err != nil {
		return []string{}, errors.New("Unable to read directory " + dirPath)
	}
	// Read the contents of the directory and return an error if it fails.

	result := make([]string, len(files))
	for i := range files {
		result[i] = strings.Split(files[i].Name(), ".")[0]
		// result[i] stores the file name without the extension.
	}
	// Loop through the files and extract the file names.

	return result, nil
}

// scandir is a function that reads the contents of a directory and returns a list of file names.
