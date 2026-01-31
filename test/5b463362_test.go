package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test5b463362(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./5b463362.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./5b463362.rb:::7:::Integer
./5b463362.rb:::8:::Integer
./5b463362.rb:::11:::Integer
./5b463362.rb:::12:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
