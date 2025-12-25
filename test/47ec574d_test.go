package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test47ec574d(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./47ec574d.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./47ec574d.rb:::16:::Union<Integer Float>
./47ec574d.rb:::18:::type mismatch: expected Union<Integer Float>, but got String for Integer.+`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
