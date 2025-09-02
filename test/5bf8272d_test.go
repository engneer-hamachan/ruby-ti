package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test5bf8272d(t *testing.T) {
	cmd := exec.Command("../ti", "./5bf8272d.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./5bf8272d.rb::133::Array<untyped>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
