package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test45f4813f(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./45f4813f.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./45f4813f.rb:::5:::untyped
./45f4813f.rb:::7:::untyped
./45f4813f.rb:::9:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
