package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestDa374c63(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./da374c63.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./da374c63.rb:::4:::NilClass
./da374c63.rb:::5:::NilClass
./da374c63.rb:::10:::Union<String Integer>
./da374c63.rb:::11:::Union<Integer String>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
