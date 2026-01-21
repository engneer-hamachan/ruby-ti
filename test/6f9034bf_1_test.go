package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test6f9034bf_1(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./6f9034bf_1.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./6f9034bf_1.rb:::7::: method 'hoge' is not defined`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
