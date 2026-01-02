package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestD29ae395(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./d29ae395.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./d29ae395.rb:::2:::NilClass
./d29ae395.rb:::3:::type mismatch: expected Proc, but got Integer for User.scope`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
