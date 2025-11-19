package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test7ce60cc4(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./7ce60cc4.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./7ce60cc4.rb:::9:::Array<String>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
