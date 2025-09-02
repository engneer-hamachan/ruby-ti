package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test1a91da13(t *testing.T) {
	cmd := exec.Command("../ti", "./1a91da13.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./1a91da13.rb::1::method 'test' is not defined for String"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
