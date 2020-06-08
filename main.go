package main

import (
	"os"

	"github.com/andremedeiros/loon/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
