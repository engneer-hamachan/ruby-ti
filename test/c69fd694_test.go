package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestC69fd694(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./c69fd694.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./c69fd694.rb:::7:::String
./c69fd694.rb:::8:::Integer
./c69fd694.rb:::9:::Union<Array<untyped> Integer>
./c69fd694.rb:::10:::Union<Integer Nil>
./c69fd694.rb:::11:::Nil`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
