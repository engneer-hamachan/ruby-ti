package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test570a6aa9(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./570a6aa9.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./570a6aa9.rb:::7:::String
./570a6aa9.rb:::8:::Array<untyped>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
