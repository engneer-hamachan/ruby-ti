package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestBc07defc(t *testing.T) {
	cmd := exec.Command("../ti", "./bc07defc.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./bc07defc.rb:::4:::Integer
./bc07defc.rb:::5:::Array<String Integer>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
