package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test4b6afb02(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./4b6afb02.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./4b6afb02.rb:::6:::String
./4b6afb02.rb:::8:::Union<Integer NilClass>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
