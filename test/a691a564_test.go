package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestA691a564(t *testing.T) {
	cmd := exec.Command("../ti", "./a691a564.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./a691a564.rb::8::Array<untyped>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
