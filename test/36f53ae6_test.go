package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test36f53ae6(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./36f53ae6.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./36f53ae6.rb:::16:::String
./36f53ae6.rb:::17:::Integer
./36f53ae6.rb:::20:::Union<Integer String>
./36f53ae6.rb:::21:::Union<Integer String>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
