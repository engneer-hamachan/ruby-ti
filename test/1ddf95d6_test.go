package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test1ddf95d6(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./1ddf95d6.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./1ddf95d6.rb:::12:::Union<Integer NilClass>
./1ddf95d6.rb:::15:::Union<Integer NilClass>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
