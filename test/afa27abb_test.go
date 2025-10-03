package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestAfa27abb(t *testing.T) {
	cmd := exec.Command("../ti", "./afa27abb.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./afa27abb.rb:::14:::Integer
./afa27abb.rb:::15:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
