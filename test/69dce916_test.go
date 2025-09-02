package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test69dce916(t *testing.T) {
	cmd := exec.Command("../ti", "./69dce916.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./69dce916.rb::1::too few arguments for Integer.+ expected (Union<Integer Float>) "

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
