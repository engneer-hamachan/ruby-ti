package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestC2cef890(t *testing.T) {
	cmd := exec.Command("../ti", "./c2cef890.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./c2cef890.rb::1::Array<untyped>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
