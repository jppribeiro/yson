package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// InputFile defines the structure of a file and options to convert
type InputFile struct {
	path string
}

func convertFile(file InputFile, writer chan<- map[string]string) {
}

// GetInputFile reads the provided arguments to build a InputFile
func GetInputFile() (InputFile, error) {
	if len(os.Args) < 2 {
		return InputFile{}, errors.New("specify a file to convert")
	}

	path := os.Args[1]

	return InputFile{path}, nil
}

// CheckFile checks if user has passed a file with correct extension and
// if file exists
func CheckFile(filename string) (bool, error) {
	_, err := isValidExtension(filename)

	if err != nil {
		return false, err
	}

	_, err = fileExists(filename)

	if err != nil {
		return false, err
	}

	return true, nil
}

func isValidExtension(filename string) (bool, error) {
	extension := filepath.Ext(filename)

	if extension != ".yaml" && extension != ".yml" {
		return false, fmt.Errorf("File %v is not a YAML", filename)
	}

	return true, nil
}

func fileExists(filename string) (bool, error) {
	_, err := os.Stat(filename)

	if err != nil && os.IsNotExist(err) {
		return false, fmt.Errorf("File %v does not exist", filename)
	}

	return true, nil
}
