package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test1624a20b(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./1624a20b.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./1624a20b.rb:::13:::Union<Integer untyped>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
