package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test467223b2(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./467223b2.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./467223b2.rb:::2:::Integer
./467223b2.rb:::3:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
