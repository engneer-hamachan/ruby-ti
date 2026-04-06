package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test759a0fa5(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./759a0fa5.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./759a0fa5.rb:::2:::Array<String Integer>
./759a0fa5.rb:::3:::Union<NilClass Integer>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
