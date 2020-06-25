//go:generate go-bindata -pkg catalog -prefix internal/catalog/data/ -o internal/catalog/bindata.go internal/catalog/data/...
package main

import (
	"os"

	"github.com/andremedeiros/loon/cmd"
	"github.com/andremedeiros/loon/internal/ui"
)

func main() {
	if err := cmd.Execute(); err != nil {
		ui.Error(err)
		os.Exit(1)
	}
}
