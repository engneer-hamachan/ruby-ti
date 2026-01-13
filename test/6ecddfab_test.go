package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test6ecddfab(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./6ecddfab.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./6ecddfab.rb:::5:::untyped
./6ecddfab.rb:::6:::Hash`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
