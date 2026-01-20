package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestAad9b81e(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./aad9b81e.rb", "--suggest", "--row=17")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `%new:::User.new(untyped) -> User:::`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
