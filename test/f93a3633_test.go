package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestF93a3633(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./f93a3633.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./f93a3633.rb:::20:::Array<untyped>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
