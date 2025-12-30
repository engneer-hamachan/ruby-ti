package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestBe5f2a5b(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./be5f2a5b.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./be5f2a5b.rb:::15:::type mismatch: expected Union<Integer Float>, but got String for Hoge.test"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
