package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestE3de8f97(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./e3de8f97.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./e3de8f97.rb:::5:::type mismatch: expected Integer, but got String for Test.keyword_json_test"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
