package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test7f2180ef(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./7f2180ef.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./7f2180ef.rb:::3:::Union<Integer String>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
