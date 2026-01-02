package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test7fd5c697(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./7fd5c697.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./7fd5c697.rb:::7:::type mismatch: expected Union<Integer Float>, but got NilClass for Integer.+
./7fd5c697.rb:::9:::type mismatch: expected String, but got Integer for String.+`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
