package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test49810cb7(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./49810cb7.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./49810cb7.rb:::5:::instance method '+' is not defined for Bool"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
