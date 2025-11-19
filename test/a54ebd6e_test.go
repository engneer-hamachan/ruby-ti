package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestA54ebd6e(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./a54ebd6e.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./a54ebd6e.rb:::18:::Hoge
./a54ebd6e.rb:::19:::String
./a54ebd6e.rb:::20:::Integer
./a54ebd6e.rb:::23:::too few arguments for Hoge.new expected (String, Integer)`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
