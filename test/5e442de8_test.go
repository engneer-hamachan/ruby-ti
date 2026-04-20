package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test5e442de8(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./5e442de8.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./5e442de8.rb:::1:::too few arguments for Error.peripheral_error expected (Integer, default String)`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
