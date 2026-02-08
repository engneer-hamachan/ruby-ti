package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestEe6a3fd5(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./ee6a3fd5.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./ee6a3fd5.rb:::15:::NilClass`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
