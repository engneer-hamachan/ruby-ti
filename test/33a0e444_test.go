package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test33a0e444(t *testing.T) {
	cmd := exec.Command("../ti", "./33a0e444.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./33a0e444.rb:::4:::String
./33a0e444.rb:::7:::Integer
./33a0e444.rb:::12:::Union<Integer Float>
./33a0e444.rb:::15:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
