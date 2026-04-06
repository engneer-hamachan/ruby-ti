package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test8d18a384(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./8d18a384.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./8d18a384.rb:::2:::Array<Integer>
./8d18a384.rb:::3:::Union<NilClass Integer>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
