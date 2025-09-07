package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test88880fc9(t *testing.T) {
	cmd := exec.Command("../ti", "./88880fc9.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./88880fc9.rb::6::Array<String Integer>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
