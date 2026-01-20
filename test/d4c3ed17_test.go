package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestD4c3ed17(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./d4c3ed17.rb", "--suggest", "--row=16")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `%hoge:::Fuga.hoge() -> Integer:::
%piyo:::User.piyo() -> Integer:::`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
