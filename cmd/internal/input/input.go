package input

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"yson.com/yson/cmd/internal/rescuer"
)

// FileData defines the structure of a file and options to convert
type FileData struct {
	Path       string
	Raw        bool
	RawData    []byte
	DataStruct map[string]interface{}
}

// FilePath checks the file argument passed to the command and verifies if it
// exists.
// Initializes and returns a FileData struct
func FilePath(isPipe bool) FileData {
	if len(os.Args) < 2 && !isPipe {
		rescuer.Exit(errors.New("specify a file to convert"))
	}

	raw := flag.Bool("raw", false, "Print raw string")

	flag.Parse()

	if !isPipe {
		path := flag.Arg(0)

		isValid(path)

		return FileData{path, *raw, nil, nil}
	}

	return FileData{"", *raw, nil, nil}
}

func isValid(filename string) {
	_, err := isValidExtension(filename)

	rescuer.Check(err)

	_, err = fileExists(filename)

	rescuer.Check(err)
}

func isValidExtension(filename string) (bool, error) {
	extension := filepath.Ext(filename)

	if extension != ".yaml" && extension != ".yml" {
		return false, fmt.Errorf("File %v is not a YAML", filename)
	}

	return true, nil
}

func fileExists(filename string) (bool, error) {
	_, err := os.Stat(filename)

	if err != nil && os.IsNotExist(err) {
		return false, fmt.Errorf("File %v does not exist", filename)
	}

	return true, nil
}
