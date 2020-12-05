package process

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
	"yson.com/yson/cmd/internal/input"
	"yson.com/yson/cmd/internal/rescuer"
)

// Yaml reads file contents and unmarshalls into map[interface{}]interface{}
// to be able to generate any map from any YAML structure
func Yaml(file input.FileData) string {
	file, err := readFileData(file)

	rescuer.Check(err)

	file, err = unmarshallData(file)

	rescuer.Check(err)

	json, err := convertToJSON(file)

	rescuer.Check(err)

	return json
}

func readFileData(file input.FileData) (input.FileData, error) {
	fileData, err := ioutil.ReadFile(file.Path)

	if err != nil {
		return input.FileData{}, fmt.Errorf("Error reading %v", file.Path)
	}

	file.RawData = fileData

	return file, nil
}

func unmarshallData(file input.FileData) (input.FileData, error) {
	data := make(map[string]interface{})

	err := yaml.Unmarshal(file.RawData, &data)

	if err != nil {
		return input.FileData{}, fmt.Errorf("Error parsing YAML data in %v, with error %v", file.Path, err)
	}

	file.DataStruct = data

	return file, nil
}

func convertToJSON(file input.FileData) (string, error) {
	jsonString, err := json.Marshal(file.DataStruct)

	if err != nil {
		return "", fmt.Errorf("Could not convert to json. Error: %v", err)
	}

	return string(jsonString), nil
}
