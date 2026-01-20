package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test03aa14e3(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./03aa14e3.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./03aa14e3.rb:::11::: method 'hoge' is not defined`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
