package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test62e5960c(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./62e5960c.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./62e5960c.rb:::2:::Union<Integer String>
./62e5960c.rb:::5:::String
./62e5960c.rb:::8:::Union<Array<String> Array<Integer>>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
