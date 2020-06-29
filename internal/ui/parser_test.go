package ui

import (
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		have string
		want string
		err  error
	}{
		{"Simple string", "Simple string", nil},
		{"This {bold:thing} is bold", "This \x1b[1mthing\x1b[0m is bold", nil},
		{"This {bold,underline:thing} is bold", "This \x1b[1m\x1b[4mthing\x1b[0m is bold", nil},
		{"This {bold:{underline:thing} is} not bold", "This \x1b[1m\x1b[4mthing\x1b[0m\x1b[1m is\x1b[0m not bold", nil},
		{"This {red:thing} is red", "This \x1b[31mthing\x1b[0m is red", nil},
		{"{blue:{bold:Info:} %s}\n", "\x1b[34m\x1b[1mInfo:\x1b[0m\x1b[34m %s\x1b[0m\n", nil},
		{"{blue:{bold:Info:} %s}}\n", "", ErrUnbalancedFormattingBlocks},
	}
	for _, tt := range tests {
		t.Run(tt.have, func(t *testing.T) {
			parsed, err := Parse(tt.have)
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
		MustParse("{blue:{bold:Info:} %s}")
	}
}
