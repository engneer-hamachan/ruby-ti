package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test377d9761(t *testing.T) {
	cmd := exec.Command("../ti", "./377d9761.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./377d9761.rb:::5:::Array<Integer>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
