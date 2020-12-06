package command

import (
	"fmt"
	"os"

	"yson.com/yson/cmd/internal/input"
	"yson.com/yson/cmd/internal/process"
)

// Run executes the command
func Run() {
	var reader *os.File
	isPipe := isPipe()
	fileData := input.FilePath(isPipe)

	if isPipe {
		reader = os.Stdin
	} else {
		reader = getFile(fileData.Path)
		defer reader.Close()
	}

	result := process.Yaml(fileData, reader)

	fmt.Println(result)
}

func isPipe() bool {
	fileInfo, _ := os.Stdin.Stat()

	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func getFile(filepath string) *os.File {
	fileReader, _ := os.Open(filepath)

	return fileReader
}
