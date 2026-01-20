package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestEd433add(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./ed433add.rb", "--suggest", "--row=19")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `%hoge:::Fugaa.hoge() -> Integer:::
%piyo:::User.piyo() -> Integer:::`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
