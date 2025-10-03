package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestFff42eb9(t *testing.T) {
	cmd := exec.Command("../ti", "./fff42eb9.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./fff42eb9.rb:::9:::too many arguments for Person.new`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
