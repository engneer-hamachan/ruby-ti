package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test62143a36(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./62143a36.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./62143a36.rb:::25:::Hoge.test is protect method
./62143a36.rb:::39:::Hoge.test is protect method`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
