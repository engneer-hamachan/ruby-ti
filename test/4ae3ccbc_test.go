package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test4ae3ccbc(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./4ae3ccbc.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./4ae3ccbc.rb:::1:::class 'Hoge' is not defined`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
