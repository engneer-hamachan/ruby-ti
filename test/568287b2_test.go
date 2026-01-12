package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test568287b2(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./568287b2.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./568287b2.rb:::4:::String
./568287b2.rb:::6:::Integer
./568287b2.rb:::15:::untyped
./568287b2.rb:::16:::untyped
./568287b2.rb:::17:::untyped
./568287b2.rb:::18:::String
./568287b2.rb:::19:::Unknown
./568287b2.rb:::25:::untyped
./568287b2.rb:::26:::untyped
./568287b2.rb:::35:::untyped
./568287b2.rb:::36:::untyped
./568287b2.rb:::37:::untyped
./568287b2.rb:::38:::untyped
./568287b2.rb:::40:::Integer`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
