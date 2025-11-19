package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test82e5842a(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./82e5842a.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./82e5842a.rb:::9:::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
