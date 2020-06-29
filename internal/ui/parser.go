package ui

import (
	"bytes"
	"fmt"

	"github.com/looplab/fsm"
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

// Parse parses strings with special formatting instructions that
// stylize text.
//
// Formatting instructgions should follow the format of:
// {<instr>[,instr]:<text>}

type parser struct {
	b     []byte
	out   bytes.Buffer
	pos   int
	stack [][]InstructionType
	fsm   *fsm.FSM
}

func NewParser(b []byte) *parser {
	p := &parser{b: b, out: bytes.Buffer{}, pos: 0, stack: nil}
	p.fsm = fsm.NewFSM(
		"begin",
		fsm.Events{
			{Name: "start_string_block", Src: []string{"begin", "end_formatting_block"}, Dst: "string_block"},
			{Name: "start_formatting_instructions_block", Src: []string{"string_block"}, Dst: "formatting_instructions_block"},
			{Name: "end_formatting_instructions_block", Src: []string{"formatting_instructions_block"}, Dst: "string_block"},
			{Name: "end", Src: []string{"string_block"}, Dst: "eof"},
		},
		fsm.Callbacks{},
	)
	return p
}

func (p *parser) readStringBlock() {
	fmt.Println("entering string block")

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func (p *parser) Parse() (string, error) {
	curstack := []InstructionType{}
	if err := p.fsm.Event("start_string_block"); err != nil {
		return "", err
	}
	for i := p.pos; i < len(p.b); i++ {
		switch p.fsm.Current() {
		case "string_block":
			switch p.b[i] {
			case '{':
				p.pos = i + 1
				if err := p.fsm.Event("start_formatting_instructions_block"); err != nil {
					return p.out.String(), err
				}
				break
			case '}':
				p.pos = i + 1
				p.stack = p.stack[:len(p.stack)-1]
				p.out.WriteString(Reset.String())
				for _, s := range p.stack {
					for _, i := range s {
						p.out.WriteString(i.String())
					}
				}
				break
			default:
				p.out.WriteByte(p.b[i])
			}
			break
		case "formatting_block":
			if err := p.fsm.Event("start_formatting_instructions_block"); err != nil {
				return p.out.String(), err
			}
			break
		case "formatting_instructions_block":
			if p.b[i] == ',' || p.b[i] == ':' {
				instr := string(p.b[p.pos:i])
				curstack = append(curstack, InstructionCodes[instr])
				p.pos = i + 1
			}
			if p.b[i] == ':' {
				for _, i := range curstack {
					p.out.WriteString(i.String())
				}
				p.stack = append(p.stack, curstack)
				curstack = []InstructionType{}
				if err := p.fsm.Event("end_formatting_instructions_block"); err != nil {
					return p.out.String(), err
				}
			}
		}
	}
	return p.out.String(), nil
}

func Parse(src string) string {
	p := NewParser([]byte(src))
	out, _ := p.Parse()
	return out
}
