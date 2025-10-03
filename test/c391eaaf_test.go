package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestC391eaaf(t *testing.T) {
	cmd := exec.Command("../ti", "./c391eaaf.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./c391eaaf.rb:::3:::type mismatch: expected String, but got Integer for String.+\n./c391eaaf.rb:::3:::Integer"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
