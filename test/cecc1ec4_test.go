package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestCecc1ec4(t *testing.T) {
	cmd := exec.Command("../ti", "./cecc1ec4.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./cecc1ec4.rb:::20:::Union<String Integer Array<untyped>>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
