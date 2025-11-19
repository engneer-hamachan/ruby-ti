package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test2f1cb329(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./2f1cb329.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./2f1cb329.rb:::8:::Union<Integer String>"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
