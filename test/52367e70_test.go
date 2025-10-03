package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test52367e70(t *testing.T) {
	cmd := exec.Command("../ti", "./52367e70.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./52367e70.rb:::6:::Array<Array<String> Integer>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
