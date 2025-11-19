package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test6b838ff6(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./6b838ff6.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./6b838ff6.rb:::6:::Array<String Integer>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
