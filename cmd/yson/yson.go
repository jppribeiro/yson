package main

import (
	"fmt"
)

func main() {
	start()
}

func start() {
	file, err := GetInputFile()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)

	_, err = CheckFile(file.path)

	if err != nil {
		fmt.Println(err)
	}
}
