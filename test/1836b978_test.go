package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test1836b978(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./1836b978.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./1836b978.rb:::4:::Integer
./1836b978.rb:::5:::Integer
./1836b978.rb:::9:::Integer
./1836b978.rb:::10:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
