package process

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
	"yson.com/yson/cmd/internal/rescuer"

	"yson.com/yson/cmd/internal/input"
)

// ReadJSON reads file contents and unsmarshals into map[interface{}]interface{}
func ReadJSON(file input.FileData) string {
	file, err := unmarshallJSON(file)

	rescuer.Check(err)

	yaml, err := convertToYAML(file)

	rescuer.Check(err)

	return yaml
}

func unmarshallJSON(file input.FileData) (input.FileData, error) {
	data := make(map[string]interface{})

	if err := json.Unmarshal(file.RawData, &data); err != nil {
		return input.FileData{}, fmt.Errorf("Error parsing JSON in %v, with error %v", file.Path, err)
	}

	file.DataStruct = data

	return file, nil
}

func convertToYAML(file input.FileData) (string, error) {
	var yamlData []byte
	var err error

	if yamlData, err = yaml.Marshal(file.DataStruct); err != nil {
		return "", fmt.Errorf("Could not convert to yaml. Error: %v", err)
	}

	return string(yamlData), nil
}
