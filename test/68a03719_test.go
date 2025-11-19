package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test68a03719(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./68a03719.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./68a03719.rb:::13:::Array<Integer String Array<untyped>>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
