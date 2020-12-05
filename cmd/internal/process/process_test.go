package process

import (
	"reflect"
	"testing"

	"yson.com/yson/cmd/internal/input"
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

	args := input.FileData{"test.yaml", false, []byte(yamlData), nil}

	tests := []struct {
		name    string
		args    input.FileData
		want    input.FileData
		wantErr bool
	}{
		{"With good yaml - pretty", args, input.FileData{"test.yaml", false, []byte(yamlData), wanted}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unmarshallData(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("unmarshallData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// got and tt.want will not be equal if compared directly with DeepEqual because
			// of the interface{} in map of DataStruct
			// For testing purposes assume that if tt.want["a"] is equal to got["a"]
			// test is ok; We don't have to test if the structure is ok.

			gotA := got.DataStruct["a"].(string)
			wantA := tt.want.DataStruct["a"].(string)

			isEqual := (got.Path == tt.want.Path) &&
				(reflect.DeepEqual(got.RawData, tt.want.RawData)) &&
				(reflect.DeepEqual(gotA, wantA))

			if !isEqual {
				t.Errorf("unmarshallData() = %v, want %v", got.DataStruct, tt.want.DataStruct)
			}
		})
	}
}

func TestConvertToJSON(t *testing.T) {
	type args struct {
		file input.FileData
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
		{"Converts map to string - pretty", args{input.FileData{"test.yaml", false, []byte{}, data}}, "{\n  \"a\": \"test\",\n  \"b\": [\n    1,\n    2\n  ]\n}", false},
		{"Converts map to string - raw", args{input.FileData{"test.yaml", true, []byte{}, data}}, "{\"a\":\"test\",\"b\":[1,2]}", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertToJSON(tt.args.file)
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
