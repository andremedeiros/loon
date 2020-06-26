package check

import (
	"fmt"
	"os/exec"
)

func Nix() error {
	utils := []string{"nix-shell", "nix-instantiate"}
	for _, util := range utils {
		if _, err := exec.LookPath(util); err != nil {
			return fmt.Errorf("cannot find nix utility: %s", util)
		}
	}
	return nil
}
