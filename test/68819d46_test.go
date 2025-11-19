package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test68819d46(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./68819d46.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./68819d46.rb:::12:::Array<Hoge Fuga>
./68819d46.rb:::15:::Hoge
./68819d46.rb:::16:::Fuga`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
