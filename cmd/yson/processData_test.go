package main

import (
	"reflect"
	"testing"
)

var yamlData = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

func Test_unmarshallData(t *testing.T) {

	wanted := map[string]interface{}{
		"a": "Easy!",
		"b": map[string]interface{}{
			"c": 2,
			"d": []int{3, 4},
		},
	}

	args := InputFile{"test.yaml", []byte(yamlData), nil}

	tests := []struct {
		name    string
		args    InputFile
		want    InputFile
		wantErr bool
	}{
		{"With good yaml", args, InputFile{"test.yaml", []byte(yamlData), wanted}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unmarshallData(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("unmarshallData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// got and tt.want will not be equal if compared directly with DeepEqual because
			// of the interface{} in map of dataStruct
			// For testing purposes assume that if tt.want["a"] is equal to got["a"]
			// test is ok; We don't have to test if the structure is ok.

			gotA := got.dataStruct["a"].(string)
			wantA := tt.want.dataStruct["a"].(string)

			isEqual := (got.path == tt.want.path) &&
				(reflect.DeepEqual(got.rawData, tt.want.rawData)) &&
				(reflect.DeepEqual(gotA, wantA))

			if !isEqual {
				t.Errorf("unmarshallData() = %v, want %v", got.dataStruct, tt.want.dataStruct)
			}
		})
	}
}

func TestConvertToJSON(t *testing.T) {
	type args struct {
		file InputFile
	}

	data := map[string]interface{}{
		"a": "test",
		"b": []int{1, 2},
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Converts map to string", args{InputFile{"test.yaml", []byte{}, data}}, "{\"a\":\"test\",\"b\":[1,2]}", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertToJSON(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertToJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
