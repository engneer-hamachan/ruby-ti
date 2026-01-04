package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test0cf18ff1(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./0cf18ff1.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./0cf18ff1.rb:::21:::Array<untyped>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
