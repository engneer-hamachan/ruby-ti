package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestE79cf689(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./e79cf689.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./e79cf689.rb:::2:::type mismatch: expected Union<Integer Float>, but got Union<Integer String> for Integer.+`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
