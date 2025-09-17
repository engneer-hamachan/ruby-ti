package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test6d61c189(t *testing.T) {
	cmd := exec.Command("../ti", "./6d61c189.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./6d61c189.rb::17::method 'test2' is not defined for Hoge
./6d61c189.rb::17::Hoge`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
