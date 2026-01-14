package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestAff30c08(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./aff30c08.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./aff30c08.rb:::6:::untyped
./aff30c08.rb:::7:::untyped
./aff30c08.rb:::12:::untyped
./aff30c08.rb:::13:::untyped`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
