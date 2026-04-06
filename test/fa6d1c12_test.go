package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestFa6d1c12(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./fa6d1c12.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./fa6d1c12.rb:::2:::Integer
./fa6d1c12.rb:::3:::String
./fa6d1c12.rb:::4:::Array<untyped>
./fa6d1c12.rb:::5:::Union<NilClass Integer>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
