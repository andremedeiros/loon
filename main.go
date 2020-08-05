//go:generate go-bindata -nometadata -pkg catalog -prefix internal/catalog/data/ -o internal/catalog/bindata.go internal/catalog/data/...
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/andremedeiros/loon/internal/cli"
	"github.com/andremedeiros/loon/internal/ui"

	_ "github.com/go-bindata/go-bindata"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for {
			s := <-sig
			err := fmt.Errorf("exiting because of %q signal", s)
			ui.Instance().Error(err)
			exitCode := 1
			if sysSig, ok := s.(syscall.Signal); ok {
				exitCode = int(sysSig)
			}
			cancel()
			os.Exit(exitCode + 128)
		}
	}()
	if err := cli.RunContext(ctx, os.Args); err != nil {
		ui.Instance().Error(err)
		os.Exit(1)
	}
}
