package process

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"reflect"

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
	jsonStr, err := resolveTree(file.DataStruct, file.Raw)

	if err != nil {
		return "", err
	}

	return jsonStr, nil
}

func resolveTree(data map[string]interface{}, raw bool) (string, error) {
	res := make(map[string]interface{})

	for k, v := range data {
		if interfaceVal := reflect.ValueOf(v); interfaceVal.Kind() == reflect.Map {
			innerMap := make(map[string]interface{})
			for _, key := range interfaceVal.MapKeys() {
				innerMap[fmt.Sprintf("%v", key)] = interfaceVal.MapIndex(key).Interface()
			}

			val, err := resolveTree(innerMap, raw)
			if err != nil {
				return "", err
			}

			res[fmt.Sprintf("%v", k)] = val
		} else {
			fmt.Println(v, reflect.TypeOf(v))
			res[fmt.Sprintf("%v", k)] = v
		}
	}

	if raw {
		fmt.Println(res)
		jsonData, err := json.Marshal(res)
		if err != nil {
			return "", fmt.Errorf("Could not convert raw json: %v", err)
		}

		return string(jsonData), nil
	}

	jsonData, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return "", fmt.Errorf("Could not convert json: %v", err)
	}

	return string(jsonData), nil

}
