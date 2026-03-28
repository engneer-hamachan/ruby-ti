package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test61700b68(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./61700b68.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./61700b68.rb:::2:::String
./61700b68.rb:::6:::Integer
./61700b68.rb:::9:::Array<untyped>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
