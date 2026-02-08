package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test0d92f674(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./0d92f674.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./0d92f674.rb:::15:::untyped`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
