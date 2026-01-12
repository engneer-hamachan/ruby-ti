package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test7b6e275b(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./7b6e275b.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./7b6e275b.rb:::10:::String
./7b6e275b.rb:::12:::Integer
./7b6e275b.rb:::14:::Union<Integer String>
./7b6e275b.rb:::17:::Union<Integer String>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
