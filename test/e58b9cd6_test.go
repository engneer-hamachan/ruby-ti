package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestE58b9cd6(t *testing.T) {
	cmd := exec.Command("../ti", "./e58b9cd6.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./e58b9cd6.rb::6::too few arguments for test expected (Union<String Integer>, hoge: ?Integer)"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
