package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test22605ec0(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./22605ec0.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./22605ec0.rb:::7:::Integer
./22605ec0.rb:::8:::String
./22605ec0.rb:::9:::Union<Integer String>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
