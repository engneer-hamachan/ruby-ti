package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test0081a0e2(t *testing.T) {
	cmd := exec.Command("../ti", "./0081a0e2.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./0081a0e2.rb::9::too few arguments for Test.test expected (Integer, String)\n./0081a0e2.rb::9::type mismatch: expected String, but got Integer for String.+"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
