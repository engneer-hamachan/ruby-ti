package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestA080b8f5(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./a080b8f5.rb", "--suggest", "--row=16")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `%f:::Fuga.f(Integer) -> Integer:::
%y:::Piyo.y() -> Integer:::`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
