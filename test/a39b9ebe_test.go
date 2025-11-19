package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestA39b9ebe(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./a39b9ebe.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./a39b9ebe.rb:::38:::Hoge.test is protect method`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
