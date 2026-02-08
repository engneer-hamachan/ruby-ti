package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestA117a6e0(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./a117a6e0.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./a117a6e0.rb:::12:::Integer
./a117a6e0.rb:::15:::NilClass`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
