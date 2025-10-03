package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestBbe9c329(t *testing.T) {
	cmd := exec.Command("../ti", "./bbe9c329.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./bbe9c329.rb:::5:::Union<Integer String Array<untyped>>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
