package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test6ce2cc88(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./6ce2cc88.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./6ce2cc88.rb:::4:::Integer
./6ce2cc88.rb:::5:::Array<String>
./6ce2cc88.rb:::6:::String
./6ce2cc88.rb:::7:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
