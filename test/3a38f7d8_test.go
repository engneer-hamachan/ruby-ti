package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test3a38f7d8(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./3a38f7d8.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./3a38f7d8.rb:::20:::Array<untyped>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
