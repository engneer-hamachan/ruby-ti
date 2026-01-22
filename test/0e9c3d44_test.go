package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test0e9c3d44(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./0e9c3d44.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./0e9c3d44.rb:::12:::class method 'yyy' is not defined for User
./0e9c3d44.rb:::12:::Unknown`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
