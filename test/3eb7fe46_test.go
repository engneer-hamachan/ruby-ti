package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test3eb7fe46(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./3eb7fe46.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./3eb7fe46.rb:::24:::String
./3eb7fe46.rb:::25:::Array<untyped>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
