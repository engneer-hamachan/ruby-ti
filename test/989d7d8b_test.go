package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test989d7d8b(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./989d7d8b.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./989d7d8b.rb:::6:::Union<Integer NilClass>
./989d7d8b.rb:::8:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
