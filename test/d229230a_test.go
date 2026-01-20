package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestD229230a(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./d229230a.rb", "--suggest", "--row=14")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `%hoge:::User.hoge() -> Integer:::`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
