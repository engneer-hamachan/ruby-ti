package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test32b29618(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./32b29618.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./32b29618.rb:::1:::Array<String Nil>
./32b29618.rb:::2:::Array<String Nil>
./32b29618.rb:::3:::Array<String Nil>
./32b29618.rb:::4:::Array<String Nil>
./32b29618.rb:::5:::Array<String Nil>
./32b29618.rb:::6:::Array<String Nil>
./32b29618.rb:::7:::Array<String Nil>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
