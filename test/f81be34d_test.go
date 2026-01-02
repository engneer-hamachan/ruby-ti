package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestF81be34d(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./f81be34d.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./f81be34d.rb:::5:::Union<NilClass String Array<untyped>>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
