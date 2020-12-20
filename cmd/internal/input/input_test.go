package input

import (
	"flag"
	"os"
	"reflect"
	"testing"
)

func Test_FilePath(t *testing.T) {
	var content = `
a: Easy!
b:
  c: 2
  d: [3, 4]
  e:
    f: 5
    g:
      - h: 6
      - 7
      - - 8
        - 9
`
	temp, _ := os.Create("test.yml")

	temp.Write([]byte(content))

	temp.Close()

	defer os.Remove("test.yml")

	tests := []struct {
		name    string
		want    FileData
		wantErr bool
		osArgs  []string
	}{
		{"Default arguments with a .yml", FileData{InputYaml, "test.yml", false, []byte(content), nil}, false, []string{"cmd", "test.yml"}},
		{"With <raw> flag", FileData{InputYaml, "test.yml", true, []byte(content), nil}, false, []string{"cmd", "--raw", "test.yml"}},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			osArgs := os.Args

			defer func() {
				os.Args = osArgs
				flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			}()

			os.Args = tt.osArgs

			got := FilePath()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFileData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidExtension(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"Valid extension .yaml", args{"test.yaml"}, true, false},
		{"Valid extension .yml", args{"test.yml"}, true, false},
		{"Valid extension .json", args{"test.json"}, true, false},
		{"Invalid extension .xyz", args{"test.xyz"}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := isValidExtension(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("isValidExtension() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isValidExtension() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileExists(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"Non existing file", args{"nonexisting.yml"}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fileExists(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("fileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
