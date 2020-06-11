package provider

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/andremedeiros/loon/catalog"
	"github.com/andremedeiros/loon/internal/provider/nix"
)

type Nix struct {
	Inputs []catalog.Entry

	tmp *os.File
}

func NewNix() *Nix {
	tmpfile, _ := ioutil.TempFile("", "nixderiv.nix")
	return &Nix{tmp: tmpfile}
}

func (n *Nix) String() string {
	return "Nix"
}

func (n *Nix) Add(e catalog.Entry) error {
	n.Inputs = append(n.Inputs, e)
	return nil
}

func (n *Nix) Install() error {
	deriv := &nix.Derivation{}

	for _, e := range n.Inputs {
		inputs := map[string]json.RawMessage{}
		json.Unmarshal(e.Payload, &inputs)

		for name, b := range inputs {
			p := nix.Package{Name: name}
			json.Unmarshal(b, &p)
			deriv.Packages = append(deriv.Packages, p)
		}
	}

	n.tmp.Write([]byte(deriv.Nix()))
	n.tmp.Close()

	exe := exec.Command("nix-shell", n.tmp.Name(), "--command", "true")
	exe.Stdout = os.Stdout
	exe.Stderr = os.Stderr
	exe.Run()
	return nil
}

func (n *Nix) Start(args []string) error {
	cmd := strings.Join(args, " ")
	exe := exec.Command("nix-shell", n.tmp.Name(), "--command", cmd)
	exe.Stdout = os.Stdout
	exe.Stderr = os.Stderr
	return exe.Run()
}
