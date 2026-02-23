package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestA1c53f70(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./a1c53f70.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./a1c53f70.rb:::4:::Array<Integer String>
./a1c53f70.rb:::5:::Unknown
./a1c53f70.rb:::11:::Union<Integer String>
./a1c53f70.rb:::12:::Unknown`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
