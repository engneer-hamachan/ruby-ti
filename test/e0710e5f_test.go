package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestE0710e5f(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./e0710e5f.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./e0710e5f.rb:::2:::type mismatch: expected Union<Integer Float>, but got Union<Integer String Array<untyped>> for Integer.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
