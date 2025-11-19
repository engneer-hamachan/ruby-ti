package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test04b6f3a1(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./04b6f3a1.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./04b6f3a1.rb:::4:::Integer
./04b6f3a1.rb:::9:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
