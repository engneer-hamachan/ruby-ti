package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestBb2aef1f(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./bb2aef1f.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./bb2aef1f.rb:::12:::Array<Integer Fuga>
./bb2aef1f.rb:::15:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
