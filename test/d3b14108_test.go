package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestD3b14108(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./d3b14108.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./d3b14108.rb:::6:::untyped
./d3b14108.rb:::10:::untyped
./d3b14108.rb:::14:::untyped
./d3b14108.rb:::24:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
