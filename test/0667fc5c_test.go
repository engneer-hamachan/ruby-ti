package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test0667fc5c(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./0667fc5c.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./0667fc5c.rb:::11::: method 'hoge' is not defined`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
