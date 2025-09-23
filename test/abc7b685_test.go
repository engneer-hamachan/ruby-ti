package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestAbc7b685(t *testing.T) {
	cmd := exec.Command("../ti", "./abc7b685.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./abc7b685.rb::13::class 'Arel' is not defined
./abc7b685.rb::19::type mismatch: expected Union<Integer Float>, but got String for Integer.+`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
