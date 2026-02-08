package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestB773cc51(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./b773cc51.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./b773cc51.rb:::12:::NilClass
./b773cc51.rb:::15:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
