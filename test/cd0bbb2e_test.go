package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestCd0bbb2e(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./cd0bbb2e.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./cd0bbb2e.rb:::5:::Union<Float NilClass>
./cd0bbb2e.rb:::8:::Array<Integer>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
