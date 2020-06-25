package git

import (
	"bufio"
	"bytes"
	"os/exec"
	"os/user"
	"strings"
)

func List() map[string]string {
	cmd := exec.Command("git", "config", "--list")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil
	}
	cfg := map[string]string{}
	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), "=", 2)
		if len(parts) < 2 {
			continue
		}
		cfg[parts[0]] = parts[1]
	}
	return cfg
}

func User() string {
	cfg := List()
	if user, ok := cfg["github.user"]; ok {
		return user
	}
	user, _ := user.Current()
	return user.Username
}
