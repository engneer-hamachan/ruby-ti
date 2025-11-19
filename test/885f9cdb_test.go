package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test885f9cdb(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./885f9cdb.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./885f9cdb.rb:::7:::Array<Integer Array<String>>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
