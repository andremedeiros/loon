package project

import (
	"reflect"
	"testing"
)

func TestNewFromPayload(t *testing.T) {
	t.Run("invalid payload", func(t *testing.T) {
		if _, err := NewFromPayload([]byte(`!`)); err == nil {
			t.Errorf("expected error but didn't get one")
		}
	})

	t.Run("valid payload", func(t *testing.T) {
		payload := `name: Awesome Tool
url: https://github.com/awesome/tool
environment:
  SOME: value
  SOME_OTHER: other value
`
		p, err := NewFromPayload([]byte(payload))
		if err != nil {
			t.Errorf("got %q", err)
			return
		}
		if p.Name != "Awesome Tool" {
			t.Errorf("wanted name to be 'Awesome tool', got %q", p.Name)
		}
		if p.URL != "https://github.com/awesome/tool" {
			t.Errorf("wanted URL to be something other than %q", p.URL)
		}
		env := map[string]string{
			"SOME":       "value",
			"SOME_OTHER": "other value",
		}
		if !reflect.DeepEqual(env, p.Environment) {
			t.Errorf("wanted %q, got %q", env, p.Environment)
		}
	})
}
