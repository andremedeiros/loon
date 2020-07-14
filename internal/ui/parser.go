package ui

import (
	"bytes"
	"errors"

	"github.com/dyrkin/fsm"
)

var (
	ErrUnbalancedFormattingBlocks = errors.New("unbalanced formatting blocks")
	ErrInstructionNotFound        = errors.New("instruction not found")
	ErrInfiniteLoop               = errors.New("infinite loop")
)

type InstructionTable map[string]string

// Parse parses strings with special formatting instructions that
// stylize text.
//
// Formatting instructgions should follow the format of:
// {<instr>[,instr]:<text>}

type parser struct {
	b     []byte
	out   bytes.Buffer
	pos   int
	stack [][]string
	tail  []string
	fsm   *fsm.FSM
	iters int
	err   error
}

func NewParser(b []byte, instructionCodes InstructionTable) *parser {
	p := &parser{b: b, out: bytes.Buffer{}, pos: 0, stack: nil, fsm: fsm.NewFSM()}
	p.fsm.StartWith("string_block", nil)
	p.fsm.When("string_block")(func(event *fsm.Event) *fsm.NextState {
		for i := p.pos; i < len(p.b); i++ {
			switch p.b[i] {
			case '{':
				if len(p.b) > i+1 && p.b[i+1] == '{' {
					i = i + 2
					continue
				}
				p.out.Write(p.b[p.pos:i])
				p.pos = i + 1
				return p.fsm.Goto("start_formatting_block")
			case '}':
				if len(p.b) > i+1 && p.b[i+1] == '}' {
					i = i + 2
					continue
				}
				p.out.Write(p.b[p.pos:i])
				p.pos = i + 1
				return p.fsm.Goto("end_formatting_block")
			}
		}
		p.out.Write(p.b[p.pos:len(p.b)])
		return p.fsm.Goto("done")
	})
	p.fsm.When("start_formatting_block")(func(event *fsm.Event) *fsm.NextState {
		p.tail = []string{}
		return p.fsm.Goto("start_formatting_instructions_block")
	})
	p.fsm.When("start_formatting_instructions_block")(func(event *fsm.Event) *fsm.NextState {
		for i := p.pos; i < len(p.b); i++ {
			if p.b[i] == ',' || p.b[i] == ':' {
				instr := string(p.b[p.pos:i])
				ins, ok := instructionCodes[instr]
				if !ok {
					p.err = ErrInstructionNotFound
					return p.fsm.Goto("done")
				}
				p.tail = append(p.tail, ins)
				p.out.WriteString(ins)
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
		p.tail = []string{}
		return p.fsm.Goto("string_block")
	})
	p.fsm.When("end_formatting_block")(func(event *fsm.Event) *fsm.NextState {
		if len(p.stack) == 0 {
			p.err = ErrUnbalancedFormattingBlocks
			return p.fsm.Goto("done")
		}
		p.stack = p.stack[:len(p.stack)-1]

		p.out.WriteString(instructionCodes["reset"])
		for _, s := range p.stack {
			for _, i := range s {
				p.out.WriteString(i)
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
			p.err = ErrInfiniteLoop
		}
	}
	return p.out.String(), p.err
}

func Parse(src string, instructionCodes map[string]string) (string, error) {
	return NewParser([]byte(src), instructionCodes).Parse()
}

func MustParse(src string, instructionCodes map[string]string) string {
	out, err := Parse(src, instructionCodes)
	if err != nil {
		panic(err)
	}
	return out
}
