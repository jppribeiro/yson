package command

import (
	"fmt"

	"yson.com/yson/cmd/internal/input"
	"yson.com/yson/cmd/internal/process"
)

// Run executes the command
func Run() {
	fileData := input.FilePath()

	result := process.Yaml(fileData)

	fmt.Println(result)
}
