package ui

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func ErrorWithOutput(banner string, stdout bytes.Buffer, stderr bytes.Buffer) {
	color.Set(color.FgRed)
	fmt.Println()
	fmt.Println(banner)
	if stdout.Len() > 0 {
		fmt.Println("-------------------- 8< stdout 8< --------------------")
		fmt.Println(strings.TrimSpace(stdout.String()))
	}
	if stderr.Len() > 0 {
		fmt.Println("-------------------- 8< stderr 8< --------------------")
		fmt.Println(strings.TrimSpace(stderr.String()))
	}
	fmt.Println("------------------------------------------------------")
	color.Unset()
}
