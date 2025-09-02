package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test2fed9201(t *testing.T) {
	cmd := exec.Command("../ti", "./2fed9201.rb")


	output, _ := cmd.CombinedOutput()


	expectedOutput := "./2fed9201.rb::2::type mismatch: expected String, but got Integer for String.+"


	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
