package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test2508c1da(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./2508c1da.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./2508c1da.rb:::2:::Integer
./2508c1da.rb:::9:::type mismatch: expected Union<Integer Float>, but got String for Integer.-`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
