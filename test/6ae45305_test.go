package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test6ae45305(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./6ae45305.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./6ae45305.rb:::2:::name: is not defined expected (name: Integer)\n./6ae45305.rb:::2:::too few arguments for Test.keyword_json_test2 expected (name: Integer)"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
