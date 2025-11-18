package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test03cd2d54(t *testing.T) {
	cmd := exec.Command("../ti", "./03cd2d54.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./03cd2d54.rb:::33:::Array<User>
./03cd2d54.rb:::35:::Array<Fuga>`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
