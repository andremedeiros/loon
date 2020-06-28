package version

import (
	"reflect"
	"testing"
)

func TestNewVersion(t *testing.T) {
	tests := []struct {
		have string
		want Version
	}{
		{"2.3.6", Version{2, 3, 6, ""}},
		{"1.0", Version{1, 0, 0, ""}},
		{"1.0.3.rev1234l", Version{1, 0, 3, "rev1234l"}},
	}
	for _, tt := range tests {
		t.Run(tt.have, func(t *testing.T) {
			v := New([]byte(tt.have))
			if !reflect.DeepEqual(v, tt.want) {
				t.Errorf("wanted %v, got %v\n", tt.want, v)
			}
		})
	}
}

func TestGreaterVersion(t *testing.T) {
	tests := []struct {
		test  string
		left  Version
		right Version
		want  bool
	}{
		{"1.0 > 1.1", Version{1, 0, 0, ""}, Version{1, 1, 0, ""}, false},
		{"1.2 > 1.1", Version{1, 2, 0, ""}, Version{1, 1, 0, ""}, true},
		{"1.1.1 > 1.1", Version{1, 1, 1, ""}, Version{1, 1, 0, ""}, true},
		{"1.2.0.foo > 1.2.0.bar", Version{1, 2, 0, "foo"}, Version{1, 2, 0, "bar"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			greater := tt.left.Greater(tt.right)
			if tt.want != greater {
				t.Errorf("wanted %v, got %v\n", tt.want, greater)
			}
		})
	}
}
