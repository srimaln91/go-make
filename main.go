package main

import (
	"fmt"
	"os"
	"github.com/srimaln91/go-build/util/build"
)

func main() {

	// Check OS arguments
	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--version":
			fmt.Fprintf(os.Stdout, "%s\n", build.String())
			return
		}
	}

	fmt.Println("test sample")
}
