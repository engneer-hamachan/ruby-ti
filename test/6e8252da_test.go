package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test6e8252da(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./6e8252da.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./6e8252da.rb:::2:::Hash`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
