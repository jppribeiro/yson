package main

import (
	"encoding/json"
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

// ConvertToJSON takes a map and Marshals it into a json string
func ConvertToJSON(file InputFile) (string, error) {
	jsonString, err := json.Marshal(file.dataStruct)

	if err != nil {
		return "", fmt.Errorf("Could not convert to json. Error: %v", err)
	}

	return string(jsonString), nil
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
	data := make(map[string]interface{})

	err := yaml.Unmarshal(file.rawData, &data)

	if err != nil {
		return InputFile{}, fmt.Errorf("Error parsing YAML data in %v, with error %v", file.path, err)
	}

	file.dataStruct = data

	return file, nil
}
