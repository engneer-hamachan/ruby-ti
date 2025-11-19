package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test1501d092(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./1501d092.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./1501d092.rb:::9:::Union<Integer Float>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
