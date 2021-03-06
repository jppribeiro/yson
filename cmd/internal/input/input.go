package input

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"yson.com/yson/cmd/internal/rescuer"
)

// FileData defines the structure of a file and options to convert
type FileData struct {
	Type       FileType
	Path       string
	Raw        bool
	RawData    []byte
	DataStruct map[string]interface{}
}

type FileType string

const (
	InputJSON FileType = "json"
	InputYaml FileType = "yaml"
)

// FilePath checks the file argument passed to the command and verifies if it
// exists.
// Initializes and returns a FileData struct
func FilePath() FileData {
	var reader *os.File
	isPipe := isPipe()

	if len(os.Args) < 2 && !isPipe {
		rescuer.Exit(errors.New("specify a file to convert"))
	}

	raw := flag.Bool("raw", false, "Print raw string")

	flag.Parse()

	path := ""

	if isPipe {
		reader = os.Stdin
	} else {
		path = flag.Arg(0)
		validate(path)
		reader = getFile(path)
		defer reader.Close()
	}

	rawData, err := readFileData(reader)

	rescuer.Check(err)

	fileType := resolveFileType(rawData)

	return FileData{Type: fileType, Path: path, Raw: *raw, RawData: rawData, DataStruct: nil}
}

func validate(filename string) {
	_, err := isValidExtension(filename)

	rescuer.Check(err)

	_, err = fileExists(filename)

	rescuer.Check(err)
}

func isValidExtension(filename string) (bool, error) {
	extension := filepath.Ext(filename)

	validExt := []string{".yaml", ".yml", ".json"}

	for _, ext := range validExt {
		if extension == ext {
			return true, nil
		}
	}

	return false, fmt.Errorf("File %v is not a YAML or JSON", filename)
}

func fileExists(filename string) (bool, error) {
	_, err := os.Stat(filename)

	if err != nil && os.IsNotExist(err) {
		return false, fmt.Errorf("File %v does not exist", filename)
	}

	return true, nil
}

func readFileData(r io.Reader) ([]byte, error) {
	scanner := bufio.NewScanner(r)

	output := ""

	for scanner.Scan() {
		output += scanner.Text() + "\n"
	}

	if len(output) == 0 {
		return []byte{}, fmt.Errorf("File is empty")
	}

	return []byte(output), nil
}

func resolveFileType(rawData []byte) FileType {
	// 123 = {   91 = [
	if rawData[0] == 123 || rawData[0] == 91 {
		return InputJSON
	}

	return InputYaml
}

func isPipe() bool {
	fileInfo, _ := os.Stdin.Stat()

	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func getFile(filepath string) *os.File {
	fileReader, _ := os.Open(filepath)

	return fileReader
}
