package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestAcefa800(t *testing.T) {
	cmd := exec.Command("../ti", "./acefa800.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./acefa800.rb:::9:::method 'special_ability' is not defined for Array"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
