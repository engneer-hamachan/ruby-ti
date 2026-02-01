package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test23749b9d(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./23749b9d.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./23749b9d.rb:::16:::String
./23749b9d.rb:::17:::String
./23749b9d.rb:::20:::Union<Integer String>
./23749b9d.rb:::21:::Union<Integer String>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
