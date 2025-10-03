package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestF23ee39b(t *testing.T) {
	cmd := exec.Command("../ti", "./f23ee39b.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := "./f23ee39b.rb:::2:::method '+' is not defined for KeyValue\n./f23ee39b.rb:::4:::type mismatch: expected Union<Integer Float>, but got Nil for Integer.+\n./f23ee39b.rb:::7:::q is not defined expected (Integer, Keyword, a: Keyword, b: ?, c: ?Nil)\n./f23ee39b.rb:::7:::too few arguments for test expected (Integer, Keyword, a: Keyword, b: ?, c: ?Nil)"

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
