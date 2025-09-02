package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test479850c4(t *testing.T) {
	cmd := exec.Command("../ti", "./479850c4.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./479850c4.rb::1::type mismatch: expected String, but got Integer for Test.test"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
