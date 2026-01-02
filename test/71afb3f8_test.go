package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test71afb3f8(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./71afb3f8.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./71afb3f8.rb:::4:::String
./71afb3f8.rb:::6:::NilClass
./71afb3f8.rb:::9:::Union<String NilClass>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
