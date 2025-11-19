package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestDebc0cf3(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./debc0cf3.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./debc0cf3.rb:::1:::cannot define parameter after '**' for test
./debc0cf3.rb:::2:::Integer
./debc0cf3.rb:::3:::Hash`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
