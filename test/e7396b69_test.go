package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestE7396b69(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./e7396b69.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./e7396b69.rb:::5:::String
./e7396b69.rb:::7:::Integer
./e7396b69.rb:::10:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
