package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test549711e3(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./549711e3.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./549711e3.rb:::12:::Integer
./549711e3.rb:::14:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
