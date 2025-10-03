package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestDcd5492b(t *testing.T) {
	cmd := exec.Command("../ti", "./dcd5492b.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./dcd5492b.rb:::3:::type mismatch: expected Array<untyped>, but got Integer for Array.replace\n./dcd5492b.rb:::5:::Array<untyped>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
