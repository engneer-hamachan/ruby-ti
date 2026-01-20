package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestC06eb4cb(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./c06eb4cb.rb", "--suggest", "--row=397")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `%class_method:::SelfExample.class_method() -> String:::
%new:::SelfExample.new(untyped) -> SelfExample:::`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
