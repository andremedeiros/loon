package check

import (
	"fmt"
	"os/exec"

	"github.com/andremedeiros/loon/internal/version"
)

var minVersion = version.New([]byte("2.3.6"))

func Nix() error {
	// Check utilities are present.
	utils := []string{"nix", "nix-shell", "nix-instantiate"}
	for _, util := range utils {
		if _, err := exec.LookPath(util); err != nil {
			return fmt.Errorf("cannot find nix utility: %s", util)
		}
	}
	// Check version is up to date enough.
	cmd := exec.Command("nix", "--version")
	out, _ := cmd.CombinedOutput()
	nixver := version.New(out)
	if minVersion.Greater(nixver) {
		return fmt.Errorf("nix must be at least 2.3.6 and you have %s", nixver.String())
	}
	return nil
}
