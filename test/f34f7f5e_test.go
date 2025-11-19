package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestF34f7f5e(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./f34f7f5e.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./f34f7f5e.rb:::2:::Integer
./f34f7f5e.rb:::3:::Hash
./f34f7f5e.rb:::6:::expected keyvalue argument for **kwargs parameter`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
