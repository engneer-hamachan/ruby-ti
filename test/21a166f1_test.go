package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test21a166f1(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./21a166f1.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./21a166f1.rb:::7:::Integer
./21a166f1.rb:::8:::Integer
./21a166f1.rb:::11:::Integer
./21a166f1.rb:::12:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
