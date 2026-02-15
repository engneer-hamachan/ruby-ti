package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestE1046ac7(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./e1046ac7.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./e1046ac7.rb:::2:::Integer
./e1046ac7.rb:::6:::untyped
./e1046ac7.rb:::13:::Integer
./e1046ac7.rb:::19:::untyped`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
