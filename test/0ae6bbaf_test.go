package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test0ae6bbaf(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./0ae6bbaf.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./0ae6bbaf.rb:::3:::String
./0ae6bbaf.rb:::7:::String
./0ae6bbaf.rb:::11:::String
./0ae6bbaf.rb:::15:::Integer
./0ae6bbaf.rb:::19:::Integer
./0ae6bbaf.rb:::23:::String
./0ae6bbaf.rb:::27:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
