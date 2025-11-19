package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test4886fdbd(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./4886fdbd.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./4886fdbd.rb:::20:::Union<String Integer Array<untyped>>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
