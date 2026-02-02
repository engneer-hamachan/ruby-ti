package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test0bb1dfc8_1(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./0bb1dfc8_1.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./0bb1dfc8_1.rb:::16:::Integer
./0bb1dfc8_1.rb:::17:::String
./0bb1dfc8_1.rb:::20:::Union<Integer String>
./0bb1dfc8_1.rb:::21:::Union<Integer String>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
