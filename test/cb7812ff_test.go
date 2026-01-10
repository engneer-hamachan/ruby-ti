package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestCb7812ff(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./cb7812ff.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./cb7812ff.rb:::12:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
