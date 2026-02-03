package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestFdd1fe8b(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./fdd1fe8b.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./fdd1fe8b.rb:::5:::Array<Integer NilClass>
./fdd1fe8b.rb:::8:::Union<Integer NilClass>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
