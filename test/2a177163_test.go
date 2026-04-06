package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test2a177163(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./2a177163.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./2a177163.rb:::2:::Array<Integer String>
./2a177163.rb:::3:::Union<NilClass Integer>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
