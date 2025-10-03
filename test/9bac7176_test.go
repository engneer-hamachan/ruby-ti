package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test9bac7176(t *testing.T) {
	cmd := exec.Command("../ti", "./9bac7176.rb")


	output, _ := cmd.CombinedOutput()


	expectedOutput := "./9bac7176.rb:::13:::type mismatch: expected Union<Integer Float>, but got String for Hoge.test"


	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
