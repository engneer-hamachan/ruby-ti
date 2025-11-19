package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test9359c2bf(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./9359c2bf.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./9359c2bf.rb:::8:::Array<Integer>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
