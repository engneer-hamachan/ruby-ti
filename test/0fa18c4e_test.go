package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test0fa18c4e(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./0fa18c4e.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./0fa18c4e.rb:::10:::Integer
./0fa18c4e.rb:::12:::Union<Integer NilClass>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
