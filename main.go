//go:generate go-bindata -pkg catalog -prefix internal/catalog/data/ -o internal/catalog/bindata.go internal/catalog/data/...
package main

import (
	"os"

	"github.com/andremedeiros/loon/cmd"
	"github.com/andremedeiros/loon/internal/ui/color"
)

func main() {
	ui := color.New()
	if err := cmd.Execute(ui); err != nil {
		ui.Error(err)
		os.Exit(1)
	}
}
