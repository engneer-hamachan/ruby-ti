package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test7f7207c4(t *testing.T) {
	cmd := exec.Command("../ti", "./7f7207c4.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./7f7207c4.rb:::7:::Array<Integer Array<String>>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
