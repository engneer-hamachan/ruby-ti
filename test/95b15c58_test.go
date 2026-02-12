package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test95b15c58(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./95b15c58.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./95b15c58.rb:::15:::Union<Integer NilClass>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
