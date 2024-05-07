package main

import (
	"os"

  "https://github.com/zepif/Test-service/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
