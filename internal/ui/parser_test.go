package ui

import (
	"testing"
)

var instructionCodes = map[string]string{
	"bold":      "|BOL|",
	"underline": "|UND|",
	"red":       "|RED|",
	"blue":      "|BLU|",
	"reset":     "|RES|",
}

func TestParse(t *testing.T) {
	tests := []struct {
		have string
		want string
		err  error
	}{
		{"Simple string", "Simple string", nil},
		{"This {bold:thing} is bold", "This |BOL|thing|RES| is bold", nil},
		{"This {bold,underline:thing} is bold", "This |BOL||UND|thing|RES| is bold", nil},
		{"This {bold:{underline:thing} is} not bold", "This |BOL||UND|thing|RES||BOL| is|RES| not bold", nil},
		{"This {red:thing} is red", "This |RED|thing|RES| is red", nil},
		{"{blue:{bold:Info:} %s}\n", "|BLU||BOL|Info:|RES||BLU| %s|RES|\n", nil},
		{"{blue:{bold:Info:} %s}}\n", "", ErrUnbalancedFormattingBlocks},
	}
	for _, tt := range tests {
		t.Run(tt.have, func(t *testing.T) {
			parsed, err := Parse(tt.have, instructionCodes)
			if tt.err == nil && err != nil {
				t.Errorf("did not expect error but got %q", err)
			} else if tt.err != nil && err == nil {
				t.Errorf("expected error but didnt get any")
			} else if tt.err != err {
				t.Errorf("wanted error %q but got %q", tt.err, err)
			}
			if tt.err == nil && parsed != tt.want {
				t.Errorf("wanted %q but got %q", tt.want, parsed)
			}
		})
	}
}

func BenchmarkParser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MustParse("{blue:{bold:Info:} %s}", instructionCodes)
	}
}
