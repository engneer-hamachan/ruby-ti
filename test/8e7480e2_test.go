package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test8e7480e2(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./8e7480e2.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./8e7480e2.rb:::14:::ActiveRecord::Relation
./8e7480e2.rb:::15:::User
./8e7480e2.rb:::17:::ActiveRecord::Relation
./8e7480e2.rb:::18:::User`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
