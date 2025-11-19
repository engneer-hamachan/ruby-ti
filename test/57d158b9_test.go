package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test57d158b9(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./57d158b9.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./57d158b9.rb:::4:::Union<String Integer>
./57d158b9.rb:::9:::Union<Integer String>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
