package main

import (
	"os"

	"github.com/zepif/Test-service/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
