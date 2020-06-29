package ui

import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		have string
		want string
	}{
		{"Simple string", "Simple string"},
		{"This {bold:thing} is bold", "This \x1b[1mthing\x1b[0m is bold"},
		{"This {bold,underline:thing} is bold", "This \x1b[1m\x1b[4mthing\x1b[0m is bold"},
		{"This {bold:{underline:thing} is} not bold", "This \x1b[1m\x1b[4mthing\x1b[0m\x1b[1m is\x1b[0m not bold"},
		{"This {red:thing} is red", "This \x1b[31mthing\x1b[0m is red"},
		{"{blue:{bold:Info:} %s}\n", "\x1b[34m\x1b[1mInfo:\x1b[0m\x1b[34m %s\x1b[0m\n"},
	}
	for _, tt := range tests {
		t.Run(tt.have, func(t *testing.T) {
			parsed, err := Parse(tt.have)
			if err != nil {
				t.Errorf("got error %w", err)
			}
			if parsed != tt.want {
				t.Errorf("wanted %q but got %q", tt.want, parsed)
			}
		})
	}
}
