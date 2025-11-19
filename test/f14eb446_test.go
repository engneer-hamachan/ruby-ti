package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestF14eb446(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./f14eb446.rb")


	output, _ := cmd.CombinedOutput()


	expectedOutput := "./f14eb446.rb:::11:::type mismatch: expected String, but got Integer for hoge"


	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
