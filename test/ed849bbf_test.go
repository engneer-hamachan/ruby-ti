package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestEd849bbf(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./ed849bbf.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./ed849bbf.rb:::4:::Integer
./ed849bbf.rb:::12::: method 'hoge' is not defined`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
