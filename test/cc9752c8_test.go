package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestCc9752c8(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./cc9752c8.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./cc9752c8.rb:::4:::Integer
./cc9752c8.rb:::5:::Array<String>
./cc9752c8.rb:::6:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
