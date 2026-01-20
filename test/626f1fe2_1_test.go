package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test626f1fe2_1(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./626f1fe2_1.rb", "--hover", "--row=12")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `%<=:::Integer.<=(Union<Integer Float>) -> Bool:::Perform less than or equal comparison
./626f1fe2_1.rb:::130:::Hash`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
