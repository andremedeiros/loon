//go:generate go-bindata -pkg catalog -prefix internal/catalog/data/ -o internal/catalog/bindata.go internal/catalog/data/...
package main

import (
	"os"

	"github.com/andremedeiros/loon/internal/cli"
	"github.com/andremedeiros/loon/internal/ui"

	_ "github.com/go-bindata/go-bindata"
)

func main() {
	if err := cli.Run(os.Args); err != nil {
		ui.Instance().Error(err)
		os.Exit(1)
	}
}
