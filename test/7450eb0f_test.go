package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test7450eb0f(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./7450eb0f.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./7450eb0f.rb:::16:::String
./7450eb0f.rb:::17:::Integer
./7450eb0f.rb:::20:::Union<Integer String>
./7450eb0f.rb:::21:::Union<Integer String>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
