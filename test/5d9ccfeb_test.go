package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test5d9ccfeb(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./5d9ccfeb.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./5d9ccfeb.rb:::12:::Union<String Integer>
./5d9ccfeb.rb:::15:::NilClass`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
