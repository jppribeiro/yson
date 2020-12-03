package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	file, err := getInputFile()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)

	isValid, err := isValidExtension(file.path)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(isValid)
}

type inputFile struct {
	path string
}

func getInputFile() (inputFile, error) {
	if len(os.Args) < 2 {
		return inputFile{}, errors.New("specify a file to convert")
	}

	path := os.Args[1]

	fmt.Println(path)

	return inputFile{path}, nil
}

func isValidExtension(filename string) (bool, error) {
	extension := filepath.Ext(filename)

	fmt.Println(extension)

	if extension != ".yaml" && extension != ".yml" {
		return false, fmt.Errorf("File %v is not a YAML", filename)
	}

	return true, nil
}

func fileExists(filename string) (bool, error) {
	_, err := os.Stat(filename)

	if err != nil && os.IsNotExist(err) {
		return false, fmt.Errorf("%v does not exist.", filename)
	}

	return true, nil
}

func convertFile(file inputFile, writer chan<- map[string]string) {
}
