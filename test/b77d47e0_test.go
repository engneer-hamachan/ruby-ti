package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestB77d47e0(t *testing.T) {
	cmd := exec.Command("../ti", "./b77d47e0.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./b77d47e0.rb:::2:::type mismatch: expected Union<Integer Float>, but got Union<Integer String> for Integer.+\n./b77d47e0.rb:::3:::type mismatch: expected Integer, but got String for Test.keyword_json_test"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
