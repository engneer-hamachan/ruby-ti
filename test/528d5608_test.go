package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test528d5608(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./528d5608.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./528d5608.rb:::5:::String
./528d5608.rb:::7:::untyped
./528d5608.rb:::9:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
