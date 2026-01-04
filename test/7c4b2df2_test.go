package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test7c4b2df2(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./7c4b2df2.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./7c4b2df2.rb:::21:::Array<untyped>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
