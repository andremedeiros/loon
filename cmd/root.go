package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/andremedeiros/loon/internal/config"
)

var rootCmd = &cobra.Command{
	Use:   "loon",
	Short: "Loon is a development acceleration tool",
	Long:  `Loon is a development acceleration tool`,
}

func makeRunE(fun func(context.Context, *config.Config, *cobra.Command, []string) error) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Read()
		if err != nil {
			return err
		}

		ctx, cancel := context.WithCancel(context.Background())

		sig := make(chan os.Signal)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			for {
				s := <-sig
				fmt.Fprintf(os.Stderr, "\nExiting because of %q signal.\n", s)

				exitCode := 1
				if sysSig, ok := s.(syscall.Signal); ok {
					exitCode = int(sysSig)
				}

				cancel()
				os.Exit(exitCode + 128)
			}
		}()

		return fun(ctx, cfg, cmd, args)
	}
}

func Execute() error {
	return rootCmd.Execute()
}
