package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestEdb15682(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./edb15682.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./edb15682.rb:::2:::Array<untyped>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
