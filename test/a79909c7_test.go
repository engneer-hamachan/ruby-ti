package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestA79909c7(t *testing.T) {
	cmd := exec.Command("../ti", "./a79909c7.rb")


	output, _ := cmd.CombinedOutput()


	expectedOutput := "./a79909c7.rb::14::type mismatch: expected Union<Integer Float>, but got String for Integer.+"


	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
