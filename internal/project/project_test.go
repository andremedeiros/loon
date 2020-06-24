package project

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func fromPayload(body string) (*Project, error) {
	tmp, _ := ioutil.TempFile("", "loon.yml")
	defer os.Remove(tmp.Name())
	tmp.WriteString(body)
	tmp.Close()
	return fromPath(tmp.Name())
}

func TestNewFromPayload(t *testing.T) {
	t.Run("invalid payload", func(t *testing.T) {
		if _, err := fromPayload(`!`); err == nil {
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
		p, err := fromPayload(payload)
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

	t.Run("payload with services", func(t *testing.T) {
		payload := `name: Awesome Tool
services:
  postgresql:
    version: 12.3
  redis:
    version: 6.0.4
`

		p, err := fromPayload(payload)
		if err != nil {
			t.Errorf("got %q", err)
			return
		}
		if len(p.Services) != 2 {
			t.Errorf("wanted 2 services, got %d", len(p.Services))
		}
	})

	t.Run("payload with tasks", func(t *testing.T) {
		payload := `name: Awesome Tool
tasks:
  generate:
    description: Generates this and that
    command: generate foo
  compile:
    command: compile the thing
`
		p, err := fromPayload(payload)
		if err != nil {
			t.Errorf("got %q", err)
			return
		}
		if len(p.Tasks) != 2 {
			t.Errorf("wanted 2 tasks, got %d", len(p.Tasks))
		}
	})
}
