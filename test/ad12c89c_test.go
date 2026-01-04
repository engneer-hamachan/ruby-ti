package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestAd12c89c(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./ad12c89c.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./ad12c89c.rb:::20:::Array<untyped>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
