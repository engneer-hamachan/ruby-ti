package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestD874e88e(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./d874e88e.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./d874e88e.rb:::5:::Integer
./d874e88e.rb:::7:::untyped
./d874e88e.rb:::9:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
