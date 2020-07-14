package check

import (
	"fmt"
	"os/exec"
	"regexp"

	"github.com/hashicorp/go-version"
)

func extractNixVersion(p []byte) string {
	re := regexp.MustCompile(`(\d+.\d+.\d+)`)
	matches := re.FindStringSubmatch(string(p))
	if matches == nil {
		return ""
	} else {
		return matches[0]
	}
}

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
	nv := version.Must(version.NewVersion(extractNixVersion(out)))
	mv := version.Must(version.NewVersion("2.4.6"))
	if mv.LessThanOrEqual(nv) {
		return fmt.Errorf("nix must be at least 2.3.6 and you have %s", nv)
	}
	return nil
}
