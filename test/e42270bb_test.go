package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestE42270bb(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./e42270bb.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./e42270bb.rb:::3:::too few arguments for ActiveRecord::Base::scope expected (Symbol, Proc)
./e42270bb.rb:::4:::too few arguments for ActiveRecord::Base::has_one expected (Symbol, optional untyped)`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
