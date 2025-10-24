package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test00bf932e(t *testing.T) {
	cmd := exec.Command("../ti", "./00bf932e.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./00bf932e.rb:::15:::Unknown
./00bf932e.rb:::18:::x: is not defined expected (x: ?, y: Integer)`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
