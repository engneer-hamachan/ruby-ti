package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestB759e5d5(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./b759e5d5.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./b759e5d5.rb:::4:::Integer
./b759e5d5.rb:::5:::Array<untyped>
./b759e5d5.rb:::6:::String
./b759e5d5.rb:::7:::String
./b759e5d5.rb:::8:::Integer
./b759e5d5.rb:::9:::NilClass`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
