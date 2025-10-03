package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test2dffd3da(t *testing.T) {
	cmd := exec.Command("../ti", "./2dffd3da.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./2dffd3da.rb:::14:::String"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
