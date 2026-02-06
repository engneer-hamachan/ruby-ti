package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test32a22768(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./32a22768.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./32a22768.rb:::13:::Union<Integer Hash>
./32a22768.rb:::14:::Union<Integer Bool String NilClass>
./32a22768.rb:::15:::Array<untyped>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
