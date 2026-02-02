package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestEadb1499(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./eadb1499.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./eadb1499.rb:::16:::Integer
./eadb1499.rb:::17:::Integer
./eadb1499.rb:::20:::Union<Integer String>
./eadb1499.rb:::21:::Union<Integer String>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
