package process

import (
	"encoding/json"
	"fmt"
	"reflect"

	"gopkg.in/yaml.v2"
	"yson.com/yson/cmd/internal/input"
	"yson.com/yson/cmd/internal/rescuer"
)

// Yaml reads file contents and unmarshalls into map[interface{}]interface{}
// to be able to generate any map from any YAML structure
func Yaml(file input.FileData) string {
	file, err := unmarshallData(file)

	rescuer.Check(err)

	json, err := convertToJSON(file)

	rescuer.Check(err)

	return json
}

func unmarshallData(file input.FileData) (input.FileData, error) {
	data := make(map[interface{}]interface{})

	err := yaml.Unmarshal(file.RawData, &data)

	if err != nil {
		return input.FileData{}, fmt.Errorf("Error parsing YAML data in %v, with error %v", file.Path, err)
	}

	file.DataStruct = resolveMap(data)

	return file, nil
}

func resolveMap(data map[interface{}]interface{}) map[string]interface{} {
	res := make(map[string]interface{})

	for k, v := range data {
		interfaceVal := reflect.ValueOf(v)
		kind := interfaceVal.Kind()

		switch kind {
		case reflect.Map:
			innerMap := make(map[interface{}]interface{})
			for _, key := range interfaceVal.MapKeys() {
				innerMap[fmt.Sprintf("%v", key)] = interfaceVal.MapIndex(key).Interface()
			}

			res[fmt.Sprintf("%v", k)] = resolveMap(innerMap)
		case reflect.Slice:
			slice := interfaceVal.Slice(0, interfaceVal.Len()).Interface().([]interface{})
			res[fmt.Sprintf("%v", k)] = resolveArray(slice)
		default:
			res[fmt.Sprintf("%v", k)] = v
		}
	}

	return res
}

func resolveArray(arr []interface{}) []interface{} {
	out := make([]interface{}, len(arr))

	for i, el := range arr {
		kind := reflect.TypeOf(el).Kind()

		switch kind {
		case reflect.Map:
			mapElem := el.(map[interface{}]interface{})
			out[i] = resolveMap(mapElem)
		case reflect.Slice:
			sliceElem := el.([]interface{})
			out[i] = resolveArray(sliceElem)
		default:
			out[i] = el
		}
	}

	return out
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
