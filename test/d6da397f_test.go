package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestD6da397f(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./d6da397f.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./d6da397f.rb:::7:::Integer
./d6da397f.rb:::8:::Integer
./d6da397f.rb:::9:::Union<Integer String>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
