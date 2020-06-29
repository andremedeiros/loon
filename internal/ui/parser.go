package ui

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/dyrkin/fsm"
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
	tail  []InstructionType
	fsm   *fsm.FSM
	iters int
	err   error
}

func NewParser(b []byte) *parser {
	p := &parser{b: b, out: bytes.Buffer{}, pos: 0, stack: nil, fsm: fsm.NewFSM()}

	p.fsm.StartWith("string_block", nil)
	p.fsm.When("string_block")(func(event *fsm.Event) *fsm.NextState {
		for i := p.pos; i < len(p.b); i++ {
			switch p.b[i] {
			case '{':
				p.out.Write(p.b[p.pos:i])
				p.pos = i + 1
				return p.fsm.Goto("start_formatting_block")
			case '}':
				p.out.Write(p.b[p.pos:i])
				p.pos = i + 1
				return p.fsm.Goto("end_formatting_block")
			}
		}
		p.out.Write(p.b[p.pos:len(p.b)])
		return p.fsm.Goto("done")
	})
	p.fsm.When("start_formatting_block")(func(event *fsm.Event) *fsm.NextState {
		p.tail = []InstructionType{}
		return p.fsm.Goto("start_formatting_instructions_block")
	})
	p.fsm.When("start_formatting_instructions_block")(func(event *fsm.Event) *fsm.NextState {
		for i := p.pos; i < len(p.b); i++ {
			if p.b[i] == ',' || p.b[i] == ':' {
				ins := InstructionCodes[string(p.b[p.pos:i])]
				p.tail = append(p.tail, ins)
				p.out.WriteString(ins.String())
				p.pos = i + 1
			}
			if p.b[i] == ':' {
				return p.fsm.Goto("end_formatting_instructions_block")
			}
		}
		return nil
	})
	p.fsm.When("end_formatting_instructions_block")(func(event *fsm.Event) *fsm.NextState {
		p.stack = append(p.stack, p.tail)
		p.tail = []InstructionType{}
		return p.fsm.Goto("string_block")
	})
	p.fsm.When("end_formatting_block")(func(event *fsm.Event) *fsm.NextState {
		if len(p.stack) == 0 {
			p.err = errors.New("tried to pop from empty stack")
			return p.fsm.Goto("done")
		}
		p.stack = p.stack[:len(p.stack)-1]
		p.out.WriteString(Reset.String())
		for _, s := range p.stack {
			for _, i := range s {
				p.out.WriteString(i.String())
			}
		}
		return p.fsm.Goto("string_block")
	})
	p.fsm.When("done")(func(event *fsm.Event) *fsm.NextState { return nil })
	return p
}

func (p *parser) Parse() (string, error) {
	iterations := 0
	for p.fsm.CurrentState() != "done" {
		p.fsm.Send(nil)
		if p.err != nil {
			break
		}
		if iterations++; iterations == 1000 {
			p.err = errors.New("entered infinite loop")
		}
	}
	return p.out.String(), p.err
}

func Parse(src string) (string, error) {
	return NewParser([]byte(src)).Parse()
}

func MustParse(src string) string {
	out, err := Parse(src)
	if err != nil {
		panic(err)
	}
	return out
}
