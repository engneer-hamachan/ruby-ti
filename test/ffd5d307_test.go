package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestFfd5d307(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./ffd5d307.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./ffd5d307.rb:::5:::Array<Bool Integer>
./ffd5d307.rb:::9:::Union<Bool Array<Integer>>
./ffd5d307.rb:::15:::Array<Integer>
./ffd5d307.rb:::17:::Array<Bool Integer>
./ffd5d307.rb:::23:::Array<Array<Integer> Integer>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
