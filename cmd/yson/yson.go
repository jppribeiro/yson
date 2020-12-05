package main

import (
	"fmt"
	"os"
)

func main() {
	start()
}

func start() {
	file, err := GetInputFile()

	err = CheckFile(file.path)

	check(err)

	file, err = ProcessYAML(file)

	check(err)

	json, err := ConvertToJSON(file)

	check(err)

	fmt.Println(json)
}

func check(e error) {
	if e != nil {
		exit(e)
	}
}

func exit(e error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", e)
	os.Exit(1)
}
