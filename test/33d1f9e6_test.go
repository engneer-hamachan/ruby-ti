package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test33d1f9e6(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./33d1f9e6.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./33d1f9e6.rb:::7:::Integer
./33d1f9e6.rb:::8:::Integer
./33d1f9e6.rb:::11:::Integer
./33d1f9e6.rb:::12:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
