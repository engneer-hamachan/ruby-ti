package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test1893f79d(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./1893f79d.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./1893f79d.rb:::19:::type mismatch: expected Union<Integer Float>, but got String for Fuga.test
./1893f79d.rb:::26:::type mismatch: expected Integer, but got String for hoge`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
