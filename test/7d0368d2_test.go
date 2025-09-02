package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test7d0368d2(t *testing.T) {
	cmd := exec.Command("../ti", "./7d0368d2.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./7d0368d2.rb::10::Union<Integer String Array<untyped>>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
