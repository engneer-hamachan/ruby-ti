package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestAfaf509a(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./afaf509a.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./afaf509a.rb:::12:::untyped
./afaf509a.rb:::15:::Union<Integer NilClass>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
