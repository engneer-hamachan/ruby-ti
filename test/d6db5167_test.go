package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestD6db5167(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./d6db5167.rb", "--suggest", "--row=16")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `%y:::Piyo.y() -> unknown:::
./d6db5167.rb:::11::: method 'f' is not defined`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
