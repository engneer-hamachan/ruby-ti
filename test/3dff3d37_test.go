package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test3dff3d37(t *testing.T) {
	cmd := exec.Command("../ti", "./3dff3d37.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./3dff3d37.rb::11::Union<String Nil>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
