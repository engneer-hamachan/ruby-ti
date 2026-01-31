package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestB4c69041(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./b4c69041.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./b4c69041.rb:::5:::String
./b4c69041.rb:::6:::Integer
./b4c69041.rb:::9:::Integer
./b4c69041.rb:::10:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
