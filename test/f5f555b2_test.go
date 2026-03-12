package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestF5f555b2(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./f5f555b2.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./f5f555b2.rb:::8:::Bool
./f5f555b2.rb:::15:::Bool`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
