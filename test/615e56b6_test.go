package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test615e56b6(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./615e56b6.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./615e56b6.rb:::20:::Integer
./615e56b6.rb:::21:::String
./615e56b6.rb:::22:::untyped
./615e56b6.rb:::27:::String
./615e56b6.rb:::28:::String
./615e56b6.rb:::29:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
