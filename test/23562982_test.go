package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test23562982(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./23562982.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./23562982.rb:::5:::String
./23562982.rb:::7:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
