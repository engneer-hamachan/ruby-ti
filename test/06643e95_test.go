package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test06643e95(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./06643e95.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./06643e95.rb:::5:::String
./06643e95.rb:::6:::Integer
./06643e95.rb:::7:::Hash`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
