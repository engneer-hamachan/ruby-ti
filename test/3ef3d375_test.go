package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test3ef3d375(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./3ef3d375.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./3ef3d375.rb:::1:::Array<String>
./3ef3d375.rb:::3:::Union<Array<String> NilClass>
./3ef3d375.rb:::5:::Union<String Integer>
./3ef3d375.rb:::6:::type mismatch: expected Union<String Integer>, but got Array<untyped> for Test.compact_union
./3ef3d375.rb:::6:::Array<untyped>
./3ef3d375.rb:::8:::String
./3ef3d375.rb:::10:::String
./3ef3d375.rb:::11:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
