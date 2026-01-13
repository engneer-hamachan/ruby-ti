package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test2712be04(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./2712be04.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./2712be04.rb:::10:::String
./2712be04.rb:::12:::untyped`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
