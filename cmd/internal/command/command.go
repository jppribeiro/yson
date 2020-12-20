package command

import (
	"fmt"
	"os"

	"yson.com/yson/cmd/internal/input"
	"yson.com/yson/cmd/internal/process"
)

// Run executes the command
func Run() {


	fileData := input.FilePath(isPipe)

	result := process.Yaml(fileData, reader)

	fmt.Println(result)
}


