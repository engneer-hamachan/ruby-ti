package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test93022c0a(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./93022c0a.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./93022c0a.rb:::6:::type mismatch: expected JS::Object, but got Integer for Object.appendChild
./93022c0a.rb:::10:::type mismatch: expected String, but got JS::Object for Object.setAttribute`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
