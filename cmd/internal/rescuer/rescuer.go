package rescuer

import (
	"fmt"
	"os"
)

// Check gracefully exits if an error is detected
func Check(e error) {
	if e != nil {
		Exit(e)
	}
}

func Exit(e error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", e)
	os.Exit(1)
}
