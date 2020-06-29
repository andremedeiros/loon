package ui

import (
	"fmt"
	"io"
)

type InstructionType int

const (
	Reset InstructionType = iota
	Bold
	Faint
	Italic
	Underline
)

const (
	Black InstructionType = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	BrightBlack
	BrightRed
	BrightGreen
	BrightYellow
	BrightBlue
	BrightMagenta
	BrightCyan
	BrightWhite
)

var InstructionCodes = map[string]InstructionType{
	"reset":          Reset,
	"bold":           Bold,
	"faint":          Faint,
	"italic":         Italic,
	"underline":      Underline,
	"black":          Black,
	"red":            Red,
	"green":          Green,
	"yellow":         Yellow,
	"blue":           Blue,
	"magenta":        Magenta,
	"cyan":           Cyan,
	"white":          White,
	"bright_black":   BrightBlack,
	"bright_red":     BrightRed,
	"bright_green":   BrightGreen,
	"bright_yellow":  BrightYellow,
	"bright_blue":    BrightBlue,
	"bright_magenta": BrightMagenta,
	"bright_cyan":    BrightCyan,
	"bright_white":   BrightWhite,
}

func (i InstructionType) String() string {
	return fmt.Sprintf("\x1b[%dm", i)
}

func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	str := MustParse(fmt.Sprintf(format, a...))
	return fmt.Fprint(w, str)
}
