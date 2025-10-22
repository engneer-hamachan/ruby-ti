package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test405ad5b3(t *testing.T) {
	cmd := exec.Command("../ti", "./405ad5b3.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./405ad5b3.rb:::4:::Unknown
./405ad5b3.rb:::9:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
