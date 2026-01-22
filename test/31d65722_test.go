package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test31d65722(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./31d65722.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./31d65722.rb:::4:::Union<Integer String>
./31d65722.rb:::12::: method 'f' is not defined`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
