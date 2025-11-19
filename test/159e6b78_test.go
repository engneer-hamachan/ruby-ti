package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test159e6b78(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./159e6b78.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./159e6b78.rb:::15:::Student
./159e6b78.rb:::17:::too few arguments for Student.new expected (String, Hash)`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
