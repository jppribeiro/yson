package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// ProcessYAML reads file contents and unmarshalls into map[interface{}]interface{}
// to be able to generate any map from any YAML structure
func ProcessYAML(file InputFile) (InputFile, error) {
	file, err := readFileData(file)

	if err != nil {
		return InputFile{}, err
	}

	file, err = unmarshallData(file)

	if err != nil {
		return InputFile{}, err
	}

	return file, nil
}

func readFileData(file InputFile) (InputFile, error) {
	fileData, err := ioutil.ReadFile(file.path)

	if err != nil {
		return InputFile{}, fmt.Errorf("Error reading %v", file.path)
	}

	file.rawData = fileData

	return file, nil
}

func unmarshallData(file InputFile) (InputFile, error) {
	data := make(map[interface{}]interface{})

	err := yaml.Unmarshal(file.rawData, &data)

	if err != nil {
		return InputFile{}, fmt.Errorf("Error parsing YAML data in %v, with error %v", file.path, err)
	}

	file.dataStruct = data

	return file, nil
}
