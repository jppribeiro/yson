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

	wanted := map[interface{}]interface{}{
		"a": "Easy!",
		"b": map[interface{}]interface{}{
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unmarshallData() = %v, want %v", got, tt.want)
			}
		})
	}
}
