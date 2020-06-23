//go:generate go-bindata -pkg catalog -prefix internal/catalog/data/ -o internal/catalog/bindata.go internal/catalog/data/...
package main

import (
	"fmt"
	"os"

	"github.com/andremedeiros/loon/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
