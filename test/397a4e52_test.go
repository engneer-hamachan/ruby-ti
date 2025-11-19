package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test397a4e52(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./397a4e52.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./397a4e52.rb:::19:::Array<untyped>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
