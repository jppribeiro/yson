package process

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

	"gopkg.in/yaml.v2"
	"yson.com/yson/cmd/internal/input"
	"yson.com/yson/cmd/internal/rescuer"
)

// Yaml reads file contents and unmarshalls into map[interface{}]interface{}
// to be able to generate any map from any YAML structure
func Yaml(file input.FileData, reader io.Reader) string {
	file, err := readFileData(file, reader)

	rescuer.Check(err)

	file, err = unmarshallData(file)

	rescuer.Check(err)

	json, err := convertToJSON(file)

	rescuer.Check(err)

	return json
}

func readFileData(file input.FileData, r io.Reader) (input.FileData, error) {
	scanner := bufio.NewScanner(r)

	output := ""

	for scanner.Scan() {
		output += scanner.Text() + "\n"
	}

	file.RawData = []byte(output)

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
	var jsonString []byte
	var err error

	if file.Raw {
		jsonString, err = json.Marshal(file.DataStruct)
	} else {
		jsonString, err = json.MarshalIndent(file.DataStruct, "", "  ")
	}

	if err != nil {
		return "", fmt.Errorf("Could not convert to json. Error: %v", err)
	}

	return string(jsonString), nil
}
