package main

import (
	"fmt"

	"yson.com/yson/cmd/internal/input"
	"yson.com/yson/cmd/internal/process"
)

func main() {
	fileData := input.FilePath()

	result := process.Yaml(fileData)

	fmt.Println(result)
}
